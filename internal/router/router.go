package router

import (
	"net/http"

	"github.com/adullahkapadia/auth-service/internal/handler"
	"github.com/adullahkapadia/auth-service/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Setup(auth *handler.AuthHandler) *chi.Mux {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5500",
			"http://127.0.0.1:5500",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
		},
		ExposedHeaders: []string{
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth Service"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Post("/register", auth.Register)
	r.Post("/login", auth.Login)

	r.With(
		middleware.Auth(auth.Config),
	).Get(
		"/profile",
		auth.Profile,
	)

	return r
}