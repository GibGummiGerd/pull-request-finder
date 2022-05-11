package app

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"pull-request-finder/configuration"
	"pull-request-finder/models"
	"time"

	"gopkg.in/yaml.v3"
)

func PullRequestFlow() error {

	config, err := configuration.LoadConfiguration()
	if err != nil {
		return err
	}

	pullRequests, err := GetPullRequests(config)
	if err != nil {
		return err
	}

	currentTime := time.Now()

	output := []models.Output{}

	for _, pr := range pullRequests {
		timeCreated := pr.CreatedAt
		timeDifference := currentTime.Sub(timeCreated)
		maxTimeDifference, err := time.ParseDuration("72h")
		if err != nil {
			return err
		}

		// pull request needs to be open and not older then three days
		if timeDifference <= maxTimeDifference && pr.State == "open" {

			combinedStatus, err := GetCombinedCommitStatus(config, pr.Head.Sha)
			if err != nil {
				return err
			}

			// map information to output model
			single := models.Output{}
			single.PullRequestMessage = pr.Title
			single.PullRequestDate = pr.CreatedAt
			single.State = combinedStatus.State
			single.Statuses = combinedStatus.Statuses
			output = append(output, single)
		}

	}

	outputByte, err := yaml.Marshal(output)
	if err != nil {
		return err
	}

	err = os.WriteFile("pullRequests.yaml", outputByte, 0666)
	if err != nil {
		return err
	}

	return nil
}

// Get information about pull requests
func GetPullRequests(config configuration.Configuration) ([]models.PullRequest, error) {

	pullRequestURL := "https://api.github.com/repos/" + config.OwnerOfRepository + "/" + config.NameOfRepository + "/pulls"

	req, err := http.NewRequest("GET", pullRequestURL, nil)
	if err != nil {
		return []models.PullRequest{}, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []models.PullRequest{}, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.PullRequest{}, err
	}

	pullRequests := []models.PullRequest{}

	err = json.Unmarshal(content, &pullRequests)
	if err != nil {
		return []models.PullRequest{}, err
	}

	return pullRequests, nil
}

// Get inforamtion about the last commit of a pull request
func GetCombinedCommitStatus(config configuration.Configuration, sha string) (models.CombinedCommitStatus, error) {

	statusURL := "https://api.github.com/repos/" + config.OwnerOfRepository + "/" + config.NameOfRepository + "/commits/" + sha + "/status"

	req, err := http.NewRequest("GET", statusURL, nil)
	if err != nil {
		return models.CombinedCommitStatus{}, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.CombinedCommitStatus{}, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.CombinedCommitStatus{}, err
	}

	combinedStatus := models.CombinedCommitStatus{}

	json.Unmarshal(content, &combinedStatus)
	if err != nil {
		return models.CombinedCommitStatus{}, err
	}

	return combinedStatus, nil
}
