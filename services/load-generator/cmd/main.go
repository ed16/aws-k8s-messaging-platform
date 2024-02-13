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
	defer local.Cancel()
	api.HandleRequests()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
