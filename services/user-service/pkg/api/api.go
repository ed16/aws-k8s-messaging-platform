package api

import (
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandleRequests() {
	log.Println("Registering handlers...")
	http.HandleFunc("/create", user.CreateUser)
	http.HandleFunc("/get", user.GetUser)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Handlers registered successfully.")
}
