package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	_jwt "github.com/dgrijalva/jwt-go"
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

func (j JWT) Create(id uint, ip string) (string, error) {

	jwt := JWT{
		ID:         id,
		Authorized: true,
		ClientIp:   ip,
		CreatedAt:  time.Now().String(),
		Identifer:  uuid.New().String(),
	}

	var claims _jwt.MapClaims

	if err := json.Unmarshal(jwt.byte(), &claims); err != nil {
		return "", err
	}

	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("APP_SECRET")))
}

func get(claims *_jwt.MapClaims) (*JWT, error) {
	claimByte, _ := json.Marshal(&claims)

	var jt JWT

	if err := json.Unmarshal(claimByte, &jt); err != nil {
		return nil, err
	}

	return &jt, nil
}
func token(raw string) (*_jwt.Token, error) {
	return _jwt.Parse(raw, func(token *_jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*_jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("APP_SECRET")), nil
	})
}

func validate(token *_jwt.Token) (*_jwt.MapClaims, error) {
	claims, ok := token.Claims.(_jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func (j JWT) byte() []byte {
	b, _ := json.Marshal(&j)
	return b
}
