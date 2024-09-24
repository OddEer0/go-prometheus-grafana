package middleware

import (
	"grafana-dashboard/internal/metrics"
	"net/http"
	"strconv"
	"time"
)

type (
	writerWrapper struct {
		statusCode string
		http.ResponseWriter
	}
)

func (w *writerWrapper) WriteHeader(statusCode int) {
	w.statusCode = strconv.Itoa(statusCode)
	w.ResponseWriter.WriteHeader(statusCode)
}

func newWriterWrapper(w http.ResponseWriter) *writerWrapper {
	return &writerWrapper{
		ResponseWriter: w,
	}
}

func Metrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := newWriterWrapper(w)
		t := time.Now()
		h.ServeHTTP(ww, r)
		metrics.RequestCount.WithLabelValues(ww.statusCode, r.Method, r.URL.String()).Inc()
		metrics.RequestDuration.WithLabelValues(ww.statusCode, r.Method, r.URL.String()).Observe(time.Since(t).Seconds())
	})
}
