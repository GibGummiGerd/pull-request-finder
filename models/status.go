package models

import "time"

type CombinedCommitStatus struct {
	State      string     `json:"state"`
	Statuses   []Statuses `json:"statuses"`
	Sha        string     `json:"sha"`
	TotalCount int        `json:"total_count"`
	Repository Repository `json:"repository"`
	CommitURL  string     `json:"commit_url"`
	URL        string     `json:"url"`
}
type Statuses struct {
	URL         string    `json:"url"`
	AvatarURL   string    `json:"avatar_url"`
	ID          int       `json:"id"`
	NodeID      string    `json:"node_id"`
	State       string    `json:"state"`
	Description string    `json:"description"`
	TargetURL   string    `json:"target_url"`
	Context     string    `json:"context"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Repository struct {
	ID               int    `json:"id"`
	NodeID           string `json:"node_id"`
	Name             string `json:"name"`
	FullName         string `json:"full_name"`
	Owner            Owner  `json:"owner"`
	Private          bool   `json:"private"`
	HTMLURL          string `json:"html_url"`
	Description      string `json:"description"`
	Fork             bool   `json:"fork"`
	URL              string `json:"url"`
	ArchiveURL       string `json:"archive_url"`
	AssigneesURL     string `json:"assignees_url"`
	BlobsURL         string `json:"blobs_url"`
	BranchesURL      string `json:"branches_url"`
	CollaboratorsURL string `json:"collaborators_url"`
	CommentsURL      string `json:"comments_url"`
	CommitsURL       string `json:"commits_url"`
	CompareURL       string `json:"compare_url"`
	ContentsURL      string `json:"contents_url"`
	ContributorsURL  string `json:"contributors_url"`
	DeploymentsURL   string `json:"deployments_url"`
	DownloadsURL     string `json:"downloads_url"`
	EventsURL        string `json:"events_url"`
	ForksURL         string `json:"forks_url"`
	GitCommitsURL    string `json:"git_commits_url"`
	GitRefsURL       string `json:"git_refs_url"`
	GitTagsURL       string `json:"git_tags_url"`
	GitURL           string `json:"git_url"`
	IssueCommentURL  string `json:"issue_comment_url"`
	IssueEventsURL   string `json:"issue_events_url"`
	IssuesURL        string `json:"issues_url"`
	KeysURL          string `json:"keys_url"`
	LabelsURL        string `json:"labels_url"`
	LanguagesURL     string `json:"languages_url"`
	MergesURL        string `json:"merges_url"`
	MilestonesURL    string `json:"milestones_url"`
	NotificationsURL string `json:"notifications_url"`
	PullsURL         string `json:"pulls_url"`
	ReleasesURL      string `json:"releases_url"`
	SSHURL           string `json:"ssh_url"`
	StargazersURL    string `json:"stargazers_url"`
	StatusesURL      string `json:"statuses_url"`
	SubscribersURL   string `json:"subscribers_url"`
	SubscriptionURL  string `json:"subscription_url"`
	TagsURL          string `json:"tags_url"`
	TeamsURL         string `json:"teams_url"`
	TreesURL         string `json:"trees_url"`
	HooksURL         string `json:"hooks_url"`
}
