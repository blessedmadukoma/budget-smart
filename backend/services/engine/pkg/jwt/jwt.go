package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWT struct {
	ID         uint   `json:"id"`
	Authorized bool   `json:"authorized"`
	ClientIp   string `json:"client_ip"`
	CreatedAt  string `json:"created_at"`
	Identifer  string `json:"identifer"`
	Raw        string `json:"raw,omitempty"`
}

// Parse JWT from Authorization header
func New(r *http.Request) (*JWT, error) {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		raw := strings.Split(bearerToken, " ")[1]

		token, err := token(raw)
		if err != nil {
			return nil, err
		}

		claims, err := validate(token)
		if err != nil {
			return nil, err
		}

		jwt, err := get(claims)
		if err != nil {
			return nil, err
		}
		jwt.Raw = raw

		return jwt, nil
	}
	return nil, fmt.Errorf("no bearer token")
}

// Create a new JWT token (standalone function for token generation)
func Create(id uint, ip string, secret string, expirationInSeconds int) (string, error) {
	// Create JWT struct with initial data
	jwtData := JWT{
		ID:         id,
		Authorized: true,
		ClientIp:   ip,
		CreatedAt:  time.Now().UTC().Format(time.RFC3339),
		Identifer:  uuid.New().String(),
	}

	// Create claims directly
	claims := jwt.MapClaims{
		"id":         jwtData.ID,
		"authorized": jwtData.Authorized,
		"client_ip":  jwtData.ClientIp,
		"created_at": jwtData.CreatedAt,
		"identifer":  jwtData.Identifer,
		"exp":        time.Now().Add(time.Second * time.Duration(expirationInSeconds)).Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func get(claims *jwt.MapClaims) (*JWT, error) {
	claimByte, _ := json.Marshal(claims)

	var jt JWT

	if err := json.Unmarshal(claimByte, &jt); err != nil {
		return nil, err
	}

	return &jt, nil
}

func token(raw string) (*jwt.Token, error) {
	return jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func validate(token *jwt.Token) (*jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}
