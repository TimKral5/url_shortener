// Package jwt handles JSON web-tokens (JWT).
package jwt

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/timkral5/url_shortener/internal/log"
)

// Handler handles authentication through JSON web-tokens.
type Handler struct {
	SigningKey string
}

// NewHandler constructs a new instance of the JWT middleware.
func NewHandler(signingKey string) *Handler {
	return &Handler{
		SigningKey: signingKey,
	}
}

// GenerateToken generates a new JSON web-token.
func (middleware *Handler) GenerateToken(user string) string {
	claims := &jwt.RegisteredClaims{
		ID: "NULL",
		Issuer: "url-shortener",
		Subject: user,
		Audience: []string{},
		ExpiresAt: nil,
		NotBefore: nil,
		IssuedAt: nil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(middleware.SigningKey))
	if err != nil {
		log.Error(err)

		return ""
	}

	return signedToken
}

// Middleware is the middleware function for the server.
func (middleware *Handler) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		identity, status := middleware.parseAuthHeader(request.Header.Get("Authorization"))

		if status != http.StatusOK {
			writer.WriteHeader(status)

			return
		}

		request.Header.Set("L-Identity", identity)
		handler.ServeHTTP(writer, request)
	})
}

func (middleware *Handler) parseAuthHeader(authHeader string) (string, int) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", http.StatusOK
	}

	trimmedAuthHeader := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(
		trimmedAuthHeader,
		func(_ *jwt.Token) (any, error) {
			return []byte(middleware.SigningKey), nil
		},
		jwt.WithValidMethods([]string{ jwt.SigningMethodHS256.Alg() }),
	)
	if err != nil {
		log.Error("Failed to parse JWT token:", err)

		return "", http.StatusInternalServerError
	}

	identity, err := token.Claims.GetSubject()
	if err != nil {
		log.Error("Failed to fetch token claim:", err)

		return "", http.StatusInternalServerError
	}

	return identity, http.StatusOK
}
