package router

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"grafana-dashboard/internal/metrics"
	"net/http"
)

func NewMon() *http.ServeMux {
	mux := http.NewServeMux()
	prometheus.MustRegister(metrics.RequestCount, metrics.RequestDuration)
	mux.Handle("/metrics", promhttp.Handler())
	return mux
}
