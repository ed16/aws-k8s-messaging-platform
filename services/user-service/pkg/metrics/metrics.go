package metrics

import (
	"log"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

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
	// SystemCPUUsage System CPU usage gauge
	SystemCPUUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage_percent",
			Help: "Current system CPU usage percentage by core",
		},
		[]string{"core"},
	)
	// SystemMemoryUsage System memory usage gauge
	SystemMemoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "system_memory_usage_percent",
			Help: "Current system memory usage percentage",
		},
	)
)

// RegisterMetrics registers the metrics.
func RegisterMetrics() {
	// Register the metrics
	prometheus.MustRegister(HTTPCreateRequestsTotal)
	prometheus.MustRegister(HTTPGetRequestsTotal)
	prometheus.MustRegister(SystemCPUUsage)
	prometheus.MustRegister(SystemMemoryUsage)
}

// CollectSystemMetrics collects and updates system-level metrics.
func CollectSystemMetrics() {
	for {
		cpuPercents, err := cpu.Percent(time.Second, true)
		if err != nil {
			log.Printf("Error collecting CPU usage: %v\n", err)
			continue
		}

		for i, percent := range cpuPercents {
			SystemCPUUsage.WithLabelValues("core" + strconv.Itoa(i)).Set(percent)
		}

		memStat, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Error collecting memory usage: %v\n", err)
			continue
		}

		SystemMemoryUsage.Set(memStat.UsedPercent)

		time.Sleep(10 * time.Second) // Adjust the sleep duration as needed.
	}
}
