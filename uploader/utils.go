package uploader

import (
	"log"
	"os"
)

func getAccessToken() string {
	token := os.Getenv("STRAVA_ACCESS_TOKEN")

	if token == "" {
		log.Fatalf("Strava Access Token Not set")
	}

	return token
}
