package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/metrics"
)

type user struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

var userStore []user

// CreateUser creates the user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userStore = append(userStore, newUser)

	w.WriteHeader(http.StatusOK)

	// Increment the counter for user-service/create endpoint
	metrics.HTTPCreateRequestsTotal.Inc()
}

// GetUser get the user.
func GetUser(w http.ResponseWriter, r *http.Request) {
	var requestedUser user
	id := r.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID is not valid", http.StatusNotFound)
		return
	}
	if i < 0 || i > len(userStore)-1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	requestedUser = userStore[i]

	err = json.NewEncoder(w).Encode(requestedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("User requested: ", i, requestedUser)
	metrics.HTTPGetRequestsTotal.Inc()
}
