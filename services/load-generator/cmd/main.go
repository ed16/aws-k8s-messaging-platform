package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/api"
	local "github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/context"
)

func main() {
	local.Ctx, local.Cancel = context.WithCancel(context.Background())
	api.HandleRequests()
	defer local.Cancel()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
