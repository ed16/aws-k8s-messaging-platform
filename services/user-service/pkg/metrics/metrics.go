package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPCreateRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_create_requests_total",
			Help: "Count of all HTTP create requests",
		},
	)

	HTTPGetRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_get_requests_total",
			Help: "Count of all HTTP get requests",
		},
	)
)

func init() {
	// Register the metrics
	prometheus.MustRegister(HTTPCreateRequestsTotal)
	prometheus.MustRegister(HTTPGetRequestsTotal)
}
