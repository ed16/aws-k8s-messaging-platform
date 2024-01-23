package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

var userStore = make(map[string]User)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userStore[user.Name] = user

	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	user, ok := userStore[name]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
