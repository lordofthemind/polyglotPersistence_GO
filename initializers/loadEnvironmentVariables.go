package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(filePath string, silent bool) error {
	if filePath == "" {
		filePath = ".env"
	}

	err := godotenv.Load(filePath)
	if err != nil {
		return fmt.Errorf("error loading environment variables from %s: %w", filePath, err)
	}

	if !silent {
		fmt.Println("Environment variables loaded from", filePath)
	}
	return nil
}
