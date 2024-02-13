package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/api"
)

func main() {
	api.HandleRequests()
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
