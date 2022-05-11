package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pr-finder/configuration"
	"pr-finder/models"
	"time"
)

func Curl() error {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl \
	//   -H "Accept: application/vnd.github.v3+json" \
	//   https://api.github.com/repos/argoproj/argo-cd/commits

	config, err := configuration.LoadConfiguration()
	if err != nil {
		return err
	}

	pullRequestURL := "https://api.github.com/repos/" + config.OwnerOfRepository + "/" + config.NameOfRepository + "/pulls"

	req, err := http.NewRequest("GET", pullRequestURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	pullRequests := []models.PullRequest{}

	err = json.Unmarshal(content, &pullRequests)

	currentTime := time.Now()

	for _, pr := range pullRequests {
		timeCreated := pr.CreatedAt
		timeDifference := currentTime.Sub(timeCreated)
		maxTimeDifference, err := time.ParseDuration("72h")
		if err != nil {
			return err
		}
		if timeDifference <= maxTimeDifference {
			fmt.Printf("\n\nPull Request Message: %v, \nPull Request Time: %v\n", pr.Title, timeCreated)

			err = GetCombinedCommitStatus(config, pr.Head.Sha)
		}

	}

	// fmt.Printf("%v", string(content))

	return nil
}

func GetCombinedCommitStatus(config configuration.Configuration, sha string) error {

	statusURL := "https://api.github.com/repos/" + config.OwnerOfRepository + "/" + config.NameOfRepository + "/argo-cd/commits/" + sha + "/status"

	req2, err := http.NewRequest("GET", statusURL, nil)
	if err != nil {
		return err
	}
	req2.Header.Set("Accept", "application/vnd.github.v3+json")

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		return err
	}
	defer resp2.Body.Close()

	content2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return err
	}

	combinedStatus := models.CombinedCommitStatus{}

	json.Unmarshal(content2, &combinedStatus)
	if err != nil {
		return err
	}

	fmt.Printf("\nCombinedStatus: %+v", combinedStatus)
	return nil
}
