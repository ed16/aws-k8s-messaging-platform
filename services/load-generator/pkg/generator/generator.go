package generator

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

var (
	userServiceURL   = "http://user-service.default.svc.cluster.local:8080"
	createUsersRate  int
	getUsersRate     int
	actualCreateRate int
)

// SetCreateUsersRate sets the rate for creating users.
func SetCreateUsersRate(w http.ResponseWriter, r *http.Request) {
	var err error
	createUsersRate, err = strconv.Atoi(r.URL.Query().Get("rate"))
	if err != nil {
		http.Error(w, "Rate must be an integer", http.StatusInternalServerError)
		log.Fatalf("Rate must be an integer: %v", err)
	}

	userCreator()

	w.WriteHeader(http.StatusOK)
}

// SetGetUsersRate sets the rate at which users are generated and fetched.
// It takes an HTTP response writer and request as parameters.
// The rate is extracted from the query parameter "rate" in the request URL.
// It starts a goroutine to asynchronously execute the userGetter function.
// If an error occurs during the execution, it returns an HTTP 500 Internal Server Error
// and logs the error message.
func SetGetUsersRate(w http.ResponseWriter, r *http.Request) {
	var err error
	getUsersRate, err = strconv.Atoi(r.URL.Query().Get("rate"))
	go func() {
		err = userGetter()
		if err != nil {
			http.Error(w, "Failed to start load generator", http.StatusInternalServerError)
			log.Fatalf("Failed to start load generator: %v", err)
		}
	}()
}

func generateRandomName() string {
	names := []string{
		"Alice", "Bob", "Charlie", "Diana", "Edward", // Original names
		"Fiona", "George", "Hannah", "Ian", "Julia", // New names
		"Kevin", "Laura", "Michael", "Nina", "Oscar", // More names
		"Paula", "Quincy", "Rachel", "Steve", "Tina", // And more
		"Umar", "Violet", "William", "Xena", "Yasmin", "Zach", // Completing the list
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(names))))
	if err != nil {
		// handle error
		return ""
	}

	return names[n.Int64()]
}

func createUser(name, fullURL string) error {
	data := map[string]string{
		"name":       name,
		"created_at": time.Now().Format("2006-01-02"),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = http.Post(fullURL, "application/json", bytes.NewBuffer(jsonData)) //nolint
	return err
}

func getUserByID(id int) (*http.Response, error) {
	fullURL := userServiceURL + "/get?id=" + strconv.Itoa(id)
	return http.Get(fullURL) //nolint
}

func userCreator() {
	log.Println("Load generation rate is set to ", createUsersRate)

	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	errors := make(chan error, 1) // Buffer to prevent blocking

	go func() {
		for createUsersRate > 0 {
			name := generateRandomName()
			fullURL := userServiceURL + "/create"
			err := createUser(name, fullURL)
			if err != nil {
				errors <- err
			}
			actualCreateRate++
		}
		if createUsersRate == 0 {
			close(quit)
		}
	}()

	// Listen for the first error
	go func() {
		for err := range errors {
			log.Printf("Error creating user: %v\n", err)
		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				// Adjust the number of goroutines based on actualCreateRate
				log.Printf("Current rate: %v\r", actualCreateRate)
				actualCreateRate = 0
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func userGetter() error {
	for i := 0; i < getUsersRate; i++ {
		resp, err := getUserByID(i)
		if err != nil {
			return err
		}
		fmt.Println(resp)
	}
	return nil
}
