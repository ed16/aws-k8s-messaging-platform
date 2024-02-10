package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	local "github.com/ed16/aws-k8s-messaging-platform/services/load-generator/pkg/context"
)

var (
	userServiceURL   = "http://user-service.default.svc.cluster.local:8080"
	createUsersRate  int
	getUsersRate     int
	actualCreateRate int
)

func SetCreateUsersRate(w http.ResponseWriter, r *http.Request) {
	var err error
	createUsersRate, err = strconv.Atoi(r.URL.Query().Get("rate"))
	if err != nil {
		http.Error(w, "Rate must be an integer", http.StatusInternalServerError)
		log.Fatalf("Rate must be an integer: %v", err)
	}
	go func() {
		err = userCreator(local.Ctx)
		if err != nil {
			http.Error(w, "Failed to start load generator", http.StatusInternalServerError)
			log.Fatalf("Failed to start load generator: %v", err)
		}
	}()
}

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
	return names[rand.Intn(len(names))]
}

func createUser(name string) error {
	data := map[string]string{
		"name":       name,
		"created_at": time.Now().Format("2006-01-02"),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fullURL := userServiceURL + "/create"
	_, err = http.Post(fullURL, "application/json", bytes.NewBuffer(jsonData))
	return err
}

func getUserByID(id int) (*http.Response, error) {
	fullURL := userServiceURL + "/get?id=" + strconv.Itoa(id)
	return http.Get(fullURL)
}

func userCreator(ctx context.Context) error {
	if createUsersRate == 0 {
		log.Println("Load generation stopped as the rate is set to 0")
		return nil
	}

	// Update actualCreateRate based on the external API call
	actualCreateRate = createUsersRate

	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	errors := make(chan error, 1) // Buffer to prevent blocking

	go func() {
		for {
			select {
			case <-ticker.C:
				// Adjust the number of goroutines based on actualCreateRate
				for i := 0; i < actualCreateRate; i++ {
					go func() {
						name := generateRandomName()
						err := createUser(name)
						if err != nil {
							errors <- err
						}
					}()
				}
			case <-quit:
				ticker.Stop()
				return
			case <-ctx.Done():
				ticker.Stop()
				close(quit)
				return
			}
		}
	}()

	// Listen for the first error
	go func() {
		for err := range errors {
			log.Printf("Error creating user: %v\n", err)
			close(quit) // Stop the load generation on first error
			break
		}
	}()

	return nil
}

func adjustRate(responseTime time.Duration, ticker **time.Ticker, manageRate func() time.Duration) {
	targetResponseTime := time.Second // Adjust this based on your needs
	if responseTime > targetResponseTime && actualCreateRate > 1 {
		// Slow down, but do not decrease below 1
		actualCreateRate--
	} else if responseTime < targetResponseTime && actualCreateRate < createUsersRate {
		// Speed up, but do not exceed the target rate
		actualCreateRate++
	}
	// Adjust the ticker to the new rate
	newDuration := manageRate()
	*ticker = time.NewTicker(newDuration)
}

func userGetter() error {
	for i := 0; i < 1000; i++ {
		resp, err := getUserByID(i)
		if err != nil {
			return err
		} else {
			fmt.Println(resp)
		}

		//time.Sleep(1 * time.Second)
	}
	return nil
}
