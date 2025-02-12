package middleware

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/awakedx/task/internal/utils"
)

const (
	adminUsername = "admin"
	adminPassword = "password"
)

func AdminMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, pass, ok := r.BasicAuth()
		if ok {
			if adminUsername == username && adminPassword == pass {
				next(w, r)
				return
			}
		} else {
			utils.WriteJSONResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	}
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func LoggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("HTTP Request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)
		next.ServeHTTP(w, r)
	})
}
