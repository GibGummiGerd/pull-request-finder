package models

import "time"

type Output struct {
	PullRequestMessage string     `yaml:"PullRequestMessage"`
	PullRequestDate    time.Time  `yaml:"PullRequestDate"`
	State              string     `yaml:"StateOfPullRequest"`
	Statuses           []Statuses `yaml:"Statuses"`
}
