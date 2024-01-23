package main

import (
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/api"
	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/generator"
)

func main() {
	// Start the load generation in a separate goroutine
	go func() {
		err := generator.Start()
		if err != nil {
			log.Fatalf("Failed to start load generator: %v", err)
		}
	}()

	// Start the web server and handle API requests
	http.HandleFunc("/", api.HandleRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
