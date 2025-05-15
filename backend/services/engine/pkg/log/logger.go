package log

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (l *Logger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture request details
		entry := l.entry.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"remote": r.RemoteAddr,
		})

		// Log the incoming request
		entry.Infof("incoming request")

		// Add logger to context
		ctx := NewContext(r.Context(), l, nil)
		r = r.WithContext(ctx)

		// Wrap the ResponseWriter to capture status
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rr, r)

		// Log response status
		entry.WithField("status", rr.statusCode).Info("response completed")
	})
}

// Helper type to capture status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
