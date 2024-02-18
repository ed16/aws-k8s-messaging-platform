package api

import (
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/generator"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HandleRequests handles incoming HTTP requests.
func HandleRequests() {
	log.Println("Registering handlers...")
	http.HandleFunc("/SetCreateUsersConnections", generator.SetCreateUsersConnections)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Handlers registered successfully.")
}
