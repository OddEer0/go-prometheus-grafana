package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("req start", slog.String("method", r.Method), slog.String("url", r.URL.String()))
		t := time.Now()

		next.ServeHTTP(w, r)

		slog.Debug("req end", slog.String("method", r.Method), slog.String("url", r.URL.String()), slog.Any("time_ms", time.Since(t).Milliseconds()))
	})
}
