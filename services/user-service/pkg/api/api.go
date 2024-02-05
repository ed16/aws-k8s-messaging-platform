package api

import (
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandleRequests() {
	http.HandleFunc("/create", user.CreateUser)
	http.HandleFunc("/get", user.GetUser)
	http.Handle("/metrics", promhttp.Handler())
}
