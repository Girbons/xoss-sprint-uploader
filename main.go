package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/strava/go.strava"
)

func getAccessToken() string {
	token := os.Getenv("STRAVA_ACCESS_TOKEN")

	if token == "" {
		log.Fatalf("Strava Access Token Not set")
	}

	return token
}

func uploadFitFile(data string) {
	token := getAccessToken()
	client := strava.NewClient(token)
	service := strava.NewUploadsService(client)

	fmt.Printf("Starting Upload to strava...\n")

	upload, err := service.
		Create(strava.FileDataTypes.FIT, "", strings.NewReader(data)).
		Private().
		Do()

	if err != nil {
		if e, ok := err.(strava.Error); ok && e.Message == "Authorization Error" {
			log.Printf("Make sure your token has 'write' permissions. You'll need implement the oauth process to get one")
		}

		log.Fatal(err)
	}
	log.Printf("Upload Complete...")
	jsonForDisplay, _ := json.Marshal(upload)
	log.Printf(string(jsonForDisplay))

	log.Printf("Waiting a 5 seconds so the upload will finish (might not)")
	time.Sleep(5 * time.Second)

	uploadSummary, err := service.Get(upload.Id).Do()
	jsonForDisplay, _ = json.Marshal(uploadSummary)
	log.Printf(string(jsonForDisplay))

	log.Printf("Your new activity is id %d", uploadSummary.ActivityId)
	log.Printf("You can view it at http://www.strava.com/activities/%d", uploadSummary.ActivityId)
}

func main() {
	// The first element in os.Args is always the program name
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file path")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file", os.Args[1])
	}

	uploadFitFile(string(data))
}
