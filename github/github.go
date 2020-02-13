package github

import (
	"net/http"

	"github.com/google/go-github/v29/github"
)

// Client decorates the github.Client with newly added github actions API calls.
type Client struct {
	*github.Client
	Actions *ActionsService
}

// NewClient creates an instance of the decorated github.Client with newly added github actions API calls.
func NewClient(httpClient *http.Client) *Client {
	ghClient := github.NewClient(httpClient)
	actionsService := &ActionsService{ghClient.Actions, ghClient}
	return &Client{ghClient, actionsService}
}
