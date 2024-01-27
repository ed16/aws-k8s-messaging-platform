package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	userServiceURL    = "http://localhost:8081/get" // ?name=John%20Doe
	messageServiceURL = "http://localhost:8081/create"
)

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

	_, err = http.Post(messageServiceURL, "application/json", bytes.NewBuffer(jsonData))
	return err
}

func getUserByID(id int) (*http.Response, error) {
	fullURL := userServiceURL + "?id=" + strconv.Itoa(id)
	return http.Get(fullURL)
}

func Start() error {

	for i := 0; i < 10000; i++ {
		name := generateRandomName()

		err := createUser(name)
		if err != nil {
			return err
		}

		//time.Sleep(1 * time.Second)
	}
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
