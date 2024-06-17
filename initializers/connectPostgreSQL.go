package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PGDB  *gorm.DB
	pgErr error
)

func ConnectToPostgresql(ctx context.Context) {
	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Read connection string from environment variable
	postgresqlConnectionString := os.Getenv("POSTGRES_DOCKER_CONNECTION_URL")
	if postgresqlConnectionString == "" {
		log.Fatal("Missing environment variable POSTGRES_DOCKER_CONNECTION_URL")
	}

	// Connect to database with retries and timeout
	maxRetries := 5
	retryDelay := 5 * time.Second
	for i := 0; i < maxRetries; i++ {
		select {
		case <-ctx.Done():
			log.Fatal("Context timed out while trying to connect to database")
		default:
			PGDB, pgErr = gorm.Open(postgres.Open(postgresqlConnectionString), &gorm.Config{})
			if pgErr == nil {
				fmt.Println("Connected to database successfully")
				return // Exit loop on successful connection
			}

			log.Printf("Connection attempt %d failed: %v\n", i+1, pgErr)
			time.Sleep(retryDelay)
		}
	}

	log.Fatal("Failed to connect to Postgres after", maxRetries, "retries")
}
