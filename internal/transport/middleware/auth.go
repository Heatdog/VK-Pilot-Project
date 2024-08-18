package middleware

import (
	"log/slog"
	"net/http"
	"strings"
)

func (handler *Handler) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		full := strings.Split(r.Header.Get("Authorization"), " ")
		if len(full) != 2 {
			handler.logger.Debug("bad header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if full[0] != "Bearer" {
			handler.logger.Debug("bad token format")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := full[1]

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
