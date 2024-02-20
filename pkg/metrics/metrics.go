package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	totalRequest    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

type Metrics interface {
	IncRequestCounter(method, path, status string)
	ObserveRequestDuration(method, path, status string, duration float64)
}

var MetricLabels = []string{"method", "path", "status"}

type namespace string

const (
	NamespaceHTTP namespace = "http"
)

func NewMetrics(ns namespace) Metrics {
	totalRequest := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: fmt.Sprintf("%v_request_total", ns),
		Help: "Total number of requests",
	}, MetricLabels)

	requestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: fmt.Sprintf("%v_request_duration_seconds", ns),
		Help: "The request latencies in seconds.",
	}, MetricLabels)

	prometheus.MustRegister(totalRequest)
	prometheus.MustRegister(requestDuration)

	return &metrics{
		totalRequest:    totalRequest,
		requestDuration: requestDuration,
	}
}

func (m *metrics) IncRequestCounter(method, path, status string) {
	m.totalRequest.WithLabelValues(method, path, status).Inc()
}

func (m *metrics) ObserveRequestDuration(method, path, status string, duration float64) {
	m.requestDuration.WithLabelValues(method, path, status).Observe(duration)
}
