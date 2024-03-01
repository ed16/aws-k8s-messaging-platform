package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ed16/aws-k8s-messaging-platform/services/user-service/pkg/metrics"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDBName         = "NewMongoDB"
	mongoCollectionName = "users"
	mongoHost           = "mongo.default.svc.cluster.local:27017"
	postgresHost        = "postgres.default.svc.cluster.local"
)

type user struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB
var mongoClient *mongo.Client

// InitDB initializes the database connection.
func InitDB() {
	var err error
	// Read database credentials from environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", postgresHost, dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping the database: %v", err)
	}
	log.Println("Successfully connected to the database")
	// SQL statement to check if the "users" table exists and create it if not
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at DATE NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Printf("Failed to create the users table: %v", err)
	} else {
		log.Println("The users table was created successfully, or already exists.")
	}
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(8)
}

// InitMongoDB initializes the MongoDB connection.
func InitMongoDB() {
	// Retrieve MongoDB credentials from environment variables
	mongoUsername := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	// Ensure both username and password are present
	if mongoUsername == "" || mongoPassword == "" {
		log.Fatal("MongoDB credentials are not set in environment variables")
	}

	// Construct the MongoDB URI using the credentials and constant host address
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", mongoUsername, mongoPassword, mongoHost)

	// Connect to MongoDB using the constructed URI
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the primary
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	mongoClient = client
}

// CreateUser creates the user in both PostgreSQL and MongoDB.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// PostgreSQL insertion
	psqlQuery := `INSERT INTO users(name, created_at) VALUES($1, $2)`
	_, err = db.Exec(psqlQuery, newUser.Name, newUser.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to insert user into PostgreSQL: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// MongoDB insertion
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, bson.M{"name": newUser.Name, "created_at": newUser.CreatedAt})
	if err != nil {
		http.Error(w, "Failed to insert user into MongoDB: "+err.Error(), http.StatusInternalServerError)
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
