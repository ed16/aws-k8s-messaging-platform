package generator

import (
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
	err := createUser("John Doe")

	// Using assert.NoError to check no error is returned
	assert.NoError(t, err, "Expected no error, got error")
}
