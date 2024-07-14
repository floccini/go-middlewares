package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const AuthUserId = "middleware.auth.userID"

func WriteUnauthorized(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusUnauthorized)
	writer.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")

		if !strings.HasPrefix(authorization, "Bearer ") {
			WriteUnauthorized(writer)
			return
		}

		encodedToken := strings.TrimPrefix(authorization, "Bearer ")

		_, err := base64.StdEncoding.DecodeString(encodedToken)

		if err != nil {
			WriteUnauthorized(writer)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
