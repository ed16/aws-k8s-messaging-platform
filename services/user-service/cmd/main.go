package main

import (
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/api"
)

func main() {
	api.HandleRequests()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
