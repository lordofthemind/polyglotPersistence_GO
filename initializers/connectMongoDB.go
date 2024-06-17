package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Client
	mgErr   error
)

func ConnectToMongoDB(ctx context.Context) {
	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Read connection string from environment variable
	mongodbConnectionString := os.Getenv("MONGO_DOCKER_CONNECTION_URL")
	if mongodbConnectionString == "" {
		log.Fatal("Missing environment variable MONGO_DOCKER_CONNECTION_URL")
	}

	// Connect to database with retries and timeout
	maxRetries := 5
	retryDelay := 5 * time.Second
	for i := 0; i < maxRetries; i++ {
		select {
		case <-ctx.Done():
			log.Fatal("Context timed out while trying to connect to database")
		default:
			MongoDB, mgErr = mongo.Connect(ctx, options.Client().ApplyURI(mongodbConnectionString))
			if mgErr == nil {
				fmt.Println("Connected to MongoDB successfully")
				return // Exit loop on successful connection
			}

			log.Printf("Connection attempt %d failed: %v\n", i+1, mgErr)
			time.Sleep(retryDelay)
		}
	}

	log.Fatal("Failed to connect to MongoDB after", maxRetries, "retries")
}
