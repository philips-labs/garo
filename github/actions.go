package github

import "github.com/google/go-github/v29/github"

// ActionsService exposes the github actions apis in the Client
type ActionsService struct {
	*github.ActionsService
	client *github.Client
}
