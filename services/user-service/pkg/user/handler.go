package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/metrics"
)

type User struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

var userStore []User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userStore = append(userStore, user)

	w.WriteHeader(http.StatusOK)
	fmt.Println("User created. user.Name:", user.Name)
	// Increment the counter for user-service/create endpoint
	metrics.HTTPCreateRequestsTotal.Inc()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// name := r.URL.Query().Get("name")

	// if name != "" {
	// 	var ok bool
	// 	user, ok = userStore[name]
	// 	if !ok {
	// 		http.Error(w, "User not found", http.StatusNotFound)
	// 		return
	// 	}
	// }
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
	user = userStore[i]

	json.NewEncoder(w).Encode(user)
	fmt.Println("User requested: ", i, user)
	metrics.HTTPGetRequestsTotal.Inc()
}
