package middleware

import (
	"log/slog"
	"net/http"
)

func (handler *Handler) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		handler.logger.Debug("verify token", slog.String("token", token))

		if token == "" {
			handler.logger.Debug("token is empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, err := handler.tokenService.Validate(r.Context(), token); err != nil {
			handler.logger.Debug(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
