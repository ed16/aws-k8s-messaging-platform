package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"runtime"
	"time"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/api"
	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/metrics"
	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/user"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	api.HandleRequests()
	go func() {
		server := &http.Server{
			Addr:         ":6060",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		log.Println(server.ListenAndServe())
	}()
	go func() {
		server := &http.Server{
			Addr:         ":8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		log.Println(server.ListenAndServe())
	}()
	user.InitDB()
	user.InitMongoDB()
	metrics.RegisterMetrics()
	go metrics.CollectSystemMetrics()
	printMetricsEverySecond()
	select {}
}

func printMetricsEverySecond() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		printMetrics()
	}
}

func printMetrics() {
	// Go runtime metrics
	numGoroutines := runtime.NumGoroutine()
	// Memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	memoryUsed := memStats.Alloc / 1024 / 1024 // in MB

	fmt.Print("\rCPUs: ", runtime.NumCPU(), "CPU Usage: ", getCPUUsage(), "% Mem used: ", memoryUsed, "Mb ", "Goroutines: ", numGoroutines, "    ")
}

func getCPUUsage() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(percent[0]))
}
