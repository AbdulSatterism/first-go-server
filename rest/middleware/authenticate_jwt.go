package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.Split(header, " ")

		if len(token) != 2 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := token[1]

		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteSecret := []byte(m.cnf.JwtSecret)
		byteMessage := []byte(message)

		h := hmac.New(sha256.New, byteSecret)
		h.Write(byteMessage)

		hash := h.Sum(nil)

		newSignature := base64Encode(hash)

		if newSignature != jwtSignature {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
