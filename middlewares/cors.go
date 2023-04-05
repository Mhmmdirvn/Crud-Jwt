package middlewares

import (
	"github.com/rs/cors"
)

func CorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		ExposedHeaders: []string{"Content-Length"},
	})
}
