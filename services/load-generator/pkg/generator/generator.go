package generator

import (
	"net/http"
	"os"
	"time"
)

var (
	running bool

	userServiceURL    = os.Getenv("USER_SERVICE_URL")
	messageServiceURL = os.Getenv("MESSAGE_SERVICE_URL")

	//test
	//os.  Setenv("USER_SERVICE_URL")
)

func Start() error {
	running = true

	for running {
		_, err := http.Get(userServiceURL)
		if err != nil {
			return err
		}

		_, err = http.Get(messageServiceURL)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func Stop() {
	running = false
}
