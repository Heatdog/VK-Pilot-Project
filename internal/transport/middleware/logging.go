package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

const (
	RequestID = "request_id"
)

type RequestIDKey struct {
	id string
}

func (handler *Handler) Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := uuid.NewString()
		ctx := context.WithValue(r.Context(), RequestIDKey{id: RequestID}, id)

		handler.logRequest(id, r)

		next.ServeHTTP(w, r.WithContext(ctx))

		handler.logResponse(id, w)
	})
}

func (handler *Handler) logRequest(id string, r *http.Request) {
	handler.logger.Info("request",
		slog.String("ID", id),
		slog.String("URL", r.URL.Path),
		slog.String("method", r.Method),
		slog.String("host", r.RemoteAddr))
}

func (handler *Handler) logResponse(id string, w http.ResponseWriter) {
	handler.logger.Info("response",
		slog.String("ID", id),
		slog.Any("response", w))
}
