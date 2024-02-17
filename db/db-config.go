package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Initialize initializes the MongoDB client.
func Initialize() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")
}
