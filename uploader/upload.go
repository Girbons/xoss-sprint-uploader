package uploader

import (
	"fmt"
	"log"
	"strings"

	strava "github.com/strava/go.strava"
)

const AUTHORIZATION_ERROR string = "Authorization Error"

func UploadFitFile(data string, private bool) {
	fmt.Println(private)
	token := getAccessToken()
	client := strava.NewClient(token)
	service := strava.NewUploadsService(client)

	fmt.Printf("Starting Upload to strava...\n")

	if private {
		_, err := service.Create(strava.FileDataTypes.FIT, "", strings.NewReader(data)).Do()
		if err != nil {
			if e, ok := err.(strava.Error); ok && e.Message == AUTHORIZATION_ERROR {
				log.Printf("Make sure your token has 'write' permissions. You'll need implement the oauth process to get one")
			}

			log.Fatal(err)
		}
	} else {
		_, err := service.
			Create(strava.FileDataTypes.FIT, "", strings.NewReader(data)).
			Private().
			Do()

		if err != nil {
			if e, ok := err.(strava.Error); ok && e.Message == AUTHORIZATION_ERROR {
				log.Printf("Make sure your token has 'write' permissions. You'll need implement the oauth process to get one")
			}

			log.Fatal(err)
		}
	}

	log.Printf("Upload Complete...")
}
