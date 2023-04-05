package middlewares

import (
	"Combine-Gorm-Mux-Jwt/controllers"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid Token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != controllers.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return controllers.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}