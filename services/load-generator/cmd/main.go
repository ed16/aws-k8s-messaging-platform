package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/api"
	local "github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/context"
	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/generator"
)

func main() {
	local.Ctx, local.Cancel = context.WithCancel(context.Background())
	defer local.Cancel()

	go func() {
		server := &http.Server{
			Addr:         ":8081",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		// Set the server handler
		server.Handler = http.DefaultServeMux

		log.Fatal(server.ListenAndServe())
	}()

	api.HandleRequests()
	go generator.RunUserCreator()
	select {}
}
