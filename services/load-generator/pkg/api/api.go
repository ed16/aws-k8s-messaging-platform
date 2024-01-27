package api

import (
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/generator"
)

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/start":
		// Start the load generation in a separate goroutine
		go func() {
			err := generator.Start()
			if err != nil {
				http.Error(w, "Failed to start load generator", http.StatusInternalServerError)
				log.Fatalf("Failed to start load generator: %v", err)
			}
		}()
	case "/stop":
		//generator.Stop()
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
