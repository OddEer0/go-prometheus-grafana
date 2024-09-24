package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	FeedbackCount = prometheus.NewCounterVec(prometheus.CounterOpts{}, []string{""})
	FeedbackTotal = prometheus.NewHistogramVec(prometheus.HistogramOpts{}, []string{""})
)
