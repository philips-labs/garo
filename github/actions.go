package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

// Runner represents a self hosted runner object
type Runner struct {
	ID     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	OS     *string `json:"os,omitempty"`
	Status *string `json:"status,omitempty"`
}

// RunnerDownload represents a self hosted runner download object
type RunnerDownload struct {
	OS           *string `json:"os,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	DownloadURL  *string `json:"download_url,omitempty"`
	Filename     *string `json:"filename,omitempty"`
}

// ActionsToken represents a create/remove token for a Github actions runner
type ActionsToken struct {
	Token     *string           `json:"token,omitempty"`
	ExpiresAt *github.Timestamp `json:"expires_at,omitempty"`
}

// Client decorates the github.Client with newly added github actions API calls.
type Client struct {
	*github.Client
}

// NewClient creates an instance of the decorated github.Client with newly added github actions API calls.
func NewClient(httpClient *http.Client) *Client {
	return &Client{github.NewClient(httpClient)}
}

// ListActionsRunners Lists all self-hosted runners for a repository. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (c *Client) ListActionsRunners(ctx context.Context, owner, repo string) ([]*Runner, *github.Response, error) {
	req, err := RestRequest(http.MethodGet, "%srepos/%s/%s/actions/runners", c.BaseURL, owner, repo)
	if err != nil {
		return nil, nil, err
	}

	var runners []*Runner
	res, err := c.Do(ctx, req, &runners)
	if err != nil {
		return nil, res, err
	}
	return runners, res, nil
}

// ListActionsRunnersDownloads Lists binaries for the self-hosted runner application that you can download and run. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (c *Client) ListActionsRunnersDownloads(ctx context.Context, owner, repo string) ([]*RunnerDownload, *github.Response, error) {
	req, err := RestRequest(http.MethodGet, "%srepos/%s/%s/actions/runners/downloads", c.BaseURL, owner, repo)
	if err != nil {
		return nil, nil, err
	}

	var runners []*RunnerDownload
	res, err := c.Do(ctx, req, &runners)
	if err != nil {
		return nil, res, err
	}
	return runners, res, nil
}

// CreateActionRunnersRegistrationToken Returns a token that you can pass to the config script. The token expires after one hour. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (c *Client) CreateActionRunnersRegistrationToken(ctx context.Context, owner, repo string) (*ActionsToken, *github.Response, error) {
	return c.CreateActionRunnersTokenRequest(ctx, owner, repo, "registration")
}

// CreateActionRunnersRemoveToken Returns a token that you can pass to remove a self-hosted runner from a repository. The token expires after one hour. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (c *Client) CreateActionRunnersRemoveToken(ctx context.Context, owner, repo string) (*ActionsToken, *github.Response, error) {
	return c.CreateActionRunnersTokenRequest(ctx, owner, repo, "remove")
}

// CreateActionRunnersTokenRequest Returns a create/remove token which can be used for the config script or remove script. The token expires after one hour. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (c *Client) CreateActionRunnersTokenRequest(ctx context.Context, owner, repo, tokenType string) (*ActionsToken, *github.Response, error) {
	req, err := RestRequest(http.MethodPost, "%srepos/%s/%s/actions/runners/%s-token", c.BaseURL, owner, repo, tokenType)
	if err != nil {
		return nil, nil, err
	}

	var token ActionsToken
	res, err := c.Do(ctx, req, &token)
	if err != nil {
		return nil, res, err
	}
	return &token, res, nil
}

// DeleteActionsRunners Forces the removal of a self-hosted runner from a repository. You can use this endpoint to completely remove the runner when the machine you were using no longer exists. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (c *Client) DeleteActionsRunners(ctx context.Context, owner, repo string, runnerID int64) (*github.Response, error) {
	req, err := RestRequest(http.MethodDelete, "%srepos/%s/%s/actions/runners/%d", c.BaseURL, owner, repo, runnerID)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, nil)
}

// RestRequest creates a http.Request for the given method and format string to build the url
func RestRequest(method string, format string, args ...interface{}) (*http.Request, error) {
	url := fmt.Sprintf(format, args...)
	return http.NewRequest(method, url, nil)
}
