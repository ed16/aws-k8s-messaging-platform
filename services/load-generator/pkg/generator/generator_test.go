package generator

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCustomClient(t *testing.T) {
	client := createCustomClient()

	// Using assert.NotNil to check the client is not nil
	assert.NotNil(t, client, "Expected non-nil client, got nil")
}

func TestSetCreateUsersConnections(t *testing.T) {
	req, err := http.NewRequest("GET", "/SetCreateUsersConnections?c=0", http.NoBody)
	require.NoError(t, err, "Failed to create new HTTP request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SetCreateUsersConnections)

	handler.ServeHTTP(rr, req)

	// Using assert.Equal to check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status code")
}

func TestGenerateRandomName(t *testing.T) {
	name := generateRandomName()

	// Using assert.NotEmpty to check the name is not empty
	assert.NotEmpty(t, name, "Expected non-empty name, got empty")
}

func TestCreateUser(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the expected endpoint
		if r.URL.Path == "/create" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK) // Respond with 200 OK
		} else {
			// If the endpoint or method does not match, respond with an error
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Update the URL to use the mock server's URL
	// Note: Ensure your createUser function allows for URL override for testing purposes
	err := createUser(server.URL, "John Doe")

	// Using assert to check no error is returned
	assert.NoError(t, err, "Expected no error, got error")
}

func NewTestServerWithURL(url string, handler http.Handler) (*httptest.Server, error) {
	ts := httptest.NewUnstartedServer(handler)
	if url != "" {
		l, err := net.Listen("tcp", url)
		if err != nil {
			return nil, err
		}
		ts.Listener.Close()
		ts.Listener = l
	}
	ts.Start()
	return ts, nil
}
