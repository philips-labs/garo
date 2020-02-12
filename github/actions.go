package github

import "github.com/google/go-github/github"

// ActionsService exposes the github actions apis in the Client
type ActionsService struct {
	client *github.Client
}
