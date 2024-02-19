package generator

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	userServiceURL  = "http://user-service.default.svc.cluster.local:8080"
	customClient    = createCustomClient()
	createUsersRate atomic.Int64
	connectionsNum  atomic.Int64
)

func createCustomClient() *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
}

// SetCreateUsersConnections sets the rate at which users are created.
func SetCreateUsersConnections(w http.ResponseWriter, r *http.Request) {
	var err error
	c, err := strconv.Atoi(r.URL.Query().Get("c"))

	if err != nil {
		http.Error(w, "Number of connections must be an integer", http.StatusInternalServerError)
		log.Fatalf("Number of connections must be an integer: %v", err)
	}
	connectionsNum.Store(int64(c))

	log.Println("\nNumber of connections is set to ", c)
	w.WriteHeader(http.StatusOK)
}

// RunUserCreator runs the user creator.
func RunUserCreator() {
	var currentGoroutines atomic.Int64

	// Channel to signal goroutines to stop.
	stopChan := make(chan struct{})

	// Function to safely start new goroutines.
	startGoroutines := func(target int64) {
		for {
			current := currentGoroutines.Load()
			if current >= target {
				return
			}
			if currentGoroutines.CompareAndSwap(current, current+1) {
				go func() {
					defer currentGoroutines.Add(-1)
					for {
						select {
						case <-stopChan:
							return
						default:
							name := generateRandomName()

							err := createUser(userServiceURL, name)
							if err != nil {
								log.Printf("Error creating user: %v\n", err)
							}
						}
					}
				}()
			}
		}
	}

	go func() {
		for {
			target := connectionsNum.Load()
			current := currentGoroutines.Load()

			if current > target {
				// Too many goroutines, send stop signals.
				for i := int64(0); i < current-target; i++ {
					stopChan <- struct{}{}
				}
			} else if current < target {
				// Not enough goroutines, start more.
				startGoroutines(target)
			}

			// Check for adjustments every second.
			fmt.Printf("\r\033[KCreate user rate: %v", createUsersRate.Load())
			createUsersRate.Store(0)

			time.Sleep(time.Second)
		}
	}()
}

func generateRandomName() string {
	names := []string{
		"Alice", "Bob", "Charlie", "Diana", "Edward",
		"Fiona", "George", "Hannah", "Ian", "Julia",
		"Kevin", "Laura", "Michael", "Nina", "Oscar",
		"Paula", "Quincy", "Rachel", "Steve", "Tina",
		"Umar", "Violet", "William", "Xena", "Yasmin",
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(names))))
	if err != nil {
		return err.Error()
	}

	return names[n.Int64()]
}

func createUser(url, name string) error {
	data := map[string]string{
		"name":       name,
		"created_at": time.Now().Format("2006-01-02"),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fullURL := url + "/create"
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := customClient.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		_, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			return err
		}
		err = resp.Body.Close()
		if err != nil {
			return err
		}
	}

	createUsersRate.Add(1)
	return nil
}
