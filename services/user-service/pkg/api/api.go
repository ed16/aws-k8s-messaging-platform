package api

import (
	"net/http"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/user"
)

func HandleRequests() {
	http.HandleFunc("/create", user.CreateUser)
	http.HandleFunc("/get", user.GetUser)
}
