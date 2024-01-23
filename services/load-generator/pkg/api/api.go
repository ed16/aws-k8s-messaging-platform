package api

import (
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/generator"
)

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/start":
		err := generator.Start()
		if err != nil {
			http.Error(w, "Failed to start load generator", http.StatusInternalServerError)
		}

	case "/stop":
		generator.Stop()
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
