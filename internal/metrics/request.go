package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "request",
			Subsystem: "http",
			Name:      "http_requests_total",
			Help:      "Total number of http requests.",
		},
		[]string{"method", "path"})
)
