package main

import (
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/api"
)

func main() {

	// Start the web server and handle API requests
	http.HandleFunc("/", api.HandleRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
