package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// HTTPCreateRequestsTotal is a counter for counting all HTTP create requests.
	HTTPCreateRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_create_requests_total",
			Help: "Count of all HTTP create requests",
		},
	)

	// HTTPGetRequestsTotal is a counter for counting all HTTP get requests.
	HTTPGetRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_get_requests_total",
			Help: "Count of all HTTP get requests",
		},
	)
)

// RegisterMetrics registers the metrics.
func RegisterMetrics() {
	// Register the metrics
	prometheus.MustRegister(HTTPCreateRequestsTotal)
	prometheus.MustRegister(HTTPGetRequestsTotal)
}
