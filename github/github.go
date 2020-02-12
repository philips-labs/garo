package github

import (
	"net/http"

	"github.com/google/go-github/github"
)

// Client decorates the github.Client with newly added github actions API calls.
type Client struct {
	*github.Client
	Actions *ActionsService
}

// NewClient creates an instance of the decorated github.Client with newly added github actions API calls.
func NewClient(httpClient *http.Client) *Client {
	ghClient := github.NewClient(httpClient)
	return &Client{ghClient, &ActionsService{ghClient}}
}
