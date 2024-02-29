package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/metrics"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type user struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB

// InitDB initializes the database connection.
func InitDB() {
	var err error
	// Read database credentials from environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := "postgres.default.svc.cluster.local"
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, dbUser, dbPassword, dbName)
	fmt.Println("Connecting to the database with connection string: ", connStr)
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping the database: %v", err)
	}
	fmt.Println("Successfully connected to the database")
	// SQL statement to check if the "users" table exists and create it if not
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at DATE NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Printf("Failed to create the users table: %v", err)
	} else {
		fmt.Println("The users table was created successfully, or already exists.")
	}
}

// CreateUser creates the user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users(name, created_at) VALUES($1, $2)`
	_, err = db.Exec(query, newUser.Name, newUser.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	metrics.HTTPCreateRequestsTotal.Inc()
}

// GetUser get the user.
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var requestedUser user
	query := `SELECT name, created_at FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&requestedUser.Name, &requestedUser.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(requestedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("User requested: ", id, requestedUser)
	metrics.HTTPGetRequestsTotal.Inc()
}
