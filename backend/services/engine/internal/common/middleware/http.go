package middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	httpauth "github.com/abbot/go-http-auth"
	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/auth"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/json"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/number"
)

type HttpMiddleware struct {
	config config.Config
}

func NewHttpMiddleware(cfg config.Config) HttpMiddleware {
	return HttpMiddleware{
		config: cfg,
	}
}

func (m HttpMiddleware) BasicAuth(next http.Handler) http.Handler {
	htpasswd := httpauth.HtpasswdFileProvider("./creds/.htpasswd")
	a := httpauth.NewBasicAuthenticator("Basic Realm", htpasswd)
	realmHeader := "Basic realm=" + strconv.Quote(a.Realm)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := a.CheckAuth(r)

		if user == "" {
			w.Header().Set("WWW-Authenticate", realmHeader)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Store user in context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CORS middleware for Chi
func (m HttpMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// origin := r.Header.Get("Origin")

		// // Convert IPWhiteLists to a slice of allowed origins
		// allowedIPs := strings.Split(m.config.IPWhiteLists, ",")
		// var allowedOrigins []string
		// for _, ip := range allowedIPs {
		// 	ip = strings.TrimSpace(ip)
		// 	if ip != "" {
		// 		// Assume a default port (e.g., 3000) to create a valid origin
		// 		allowedOrigins = append(allowedOrigins, fmt.Sprintf("http://%s:3000", ip))
		// 	}
		// }

		// // Check if the origin matches any allowed origin
		// var isAllowed bool
		// for _, allowedOrigin := range allowedOrigins {
		// 	if origin == allowedOrigin {
		// 		isAllowed = true
		// 		w.Header().Set("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }

		// if !isAllowed && origin != "" {
		// 	w.WriteHeader(http.StatusForbidden)
		// 	json.WriteError(w, http.StatusForbidden, fmt.Errorf("CORS policy: Origin %s not allowed", origin))
		// 	return
		// }

		w.Header().Set("Access-Control-Allow-Origin", "*") // because my allow credentials is set to true
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Api-Version, Platform")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var singleRequestLocks sync.Map

// SingleRequest middleware for Chi
func (m HttpMiddleware) SingleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, err := auth.GetAccount(r.Context())
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, messages.ErrUnauthorized)
			return
		}

		key := number.UintToString(a.ID)

		_, loaded := singleRequestLocks.LoadOrStore(key, struct{}{})
		if loaded {
			json.WriteError(w, http.StatusBadRequest, messages.ErrBadRequest)
			return
		}

		defer singleRequestLocks.Delete(key)

		next.ServeHTTP(w, r)
	})
}

func (m HttpMiddleware) Throttle(limit uint64) func(http.Handler) http.Handler {
	// Create a map to track requests per IP
	ipRequestCounts := make(map[string]*throttleCounter)
	var mu sync.Mutex

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get client IP
			ip := getClientIP(r)

			mu.Lock()
			// Initialize counter for IP if it doesn't exist
			if _, exists := ipRequestCounts[ip]; !exists {
				ipRequestCounts[ip] = &throttleCounter{
					count:     0,
					resetTime: time.Now().Add(time.Minute),
				}
			}

			// Check if we need to reset the counter
			if time.Now().After(ipRequestCounts[ip].resetTime) {
				ipRequestCounts[ip].count = 0
				ipRequestCounts[ip].resetTime = time.Now().Add(time.Minute)
			}

			// Check if limit exceeded
			if ipRequestCounts[ip].count >= limit {
				mu.Unlock()

				json.WriteError(w, http.StatusTooManyRequests, messages.ErrThrottleTryAgainLater)
				return
			}

			// Increment count and unlock
			ipRequestCounts[ip].count++
			mu.Unlock()

			next.ServeHTTP(w, r)
		})
	}
}

func (m HttpMiddleware) EnhanceContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get client IP
        clientIP := getClientIP(r)
        
        // Create an enhanced context with both values
        ctx := r.Context()
        ctx = context.WithValue(ctx, "httpResponseWriter", w)
        ctx = context.WithValue(ctx, "clientIP", clientIP)
        
        // Continue with the enhanced context
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// IsIpWhitelisted middleware for Chi
func (m HttpMiddleware) IsIpWhitelisted(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedIPs := strings.Split(os.Getenv("WHITELISTED_IPS"), ",")

		containsIP := func(ipList []string, targetIP string) bool {
			for _, ip := range ipList {
				if ip == targetIP {
					return true
				}
			}
			return false
		}

		clientIP := getClientIP(r)

		if !containsIP(allowedIPs, clientIP) {
			errMsg := fmt.Sprintf("Access denied for IP: %s", clientIP)
			w.WriteHeader(http.StatusUnauthorized)
			json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("%s", errMsg))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Helper function to get client IP (similar to gin's ClientIP)
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		return xff[:i]
	}

	// Check X-Real-IP header
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return xrip
	}

	// Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// Helper struct for throttling
type throttleCounter struct {
	count     uint64
	resetTime time.Time
}
