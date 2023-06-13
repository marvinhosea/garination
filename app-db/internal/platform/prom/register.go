package prom

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	ResponseDuration *prometheus.SummaryVec
	RequestCount     *prometheus.CounterVec
	ResponseStatus   *prometheus.CounterVec
}

func NewMetrics() Metrics {
	var (
		requestDuration = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: "grpc_request_duration_seconds",
				Help: "Duration of gRPC requests in seconds",
			},
			[]string{"method"},
		)
		requestCount = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "grpc_request_count_total",
				Help: "Total number of gRPC requests",
			},
			[]string{"method"},
		)
		responseStatus = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "grpc_response_status",
				Help: "Response status of gRPC requests",
			},
			[]string{"method", "status"},
		)
	)

	prometheus.MustRegister(requestDuration, requestCount)

	return Metrics{
		ResponseDuration: requestDuration,
		RequestCount:     requestCount,
		ResponseStatus:   responseStatus,
	}
}
