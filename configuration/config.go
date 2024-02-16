package configurations

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	// Load env file
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading env file")
	}
	return nil
}
