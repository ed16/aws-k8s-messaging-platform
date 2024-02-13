package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/api"
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
	select {}
}
