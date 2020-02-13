package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v29/github"
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

// ListRunners Lists all self-hosted runners for a repository. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (s *ActionsService) ListRunners(ctx context.Context, owner, repo string) ([]*Runner, *github.Response, error) {
	req, err := RestRequest(http.MethodGet, "%srepos/%s/%s/actions/runners", s.client.BaseURL, owner, repo)
	if err != nil {
		return nil, nil, err
	}

	var runners []*Runner
	res, err := s.client.Do(ctx, req, &runners)
	if err != nil {
		return nil, res, err
	}
	return runners, res, nil
}

// ListRunnersDownloads Lists binaries for the self-hosted runner application that you can download and run. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (s *ActionsService) ListRunnersDownloads(ctx context.Context, owner, repo string) ([]*RunnerDownload, *github.Response, error) {
	req, err := RestRequest(http.MethodGet, "%srepos/%s/%s/actions/runners/downloads", s.client.BaseURL, owner, repo)
	if err != nil {
		return nil, nil, err
	}

	var runners []*RunnerDownload
	res, err := s.client.Do(ctx, req, &runners)
	if err != nil {
		return nil, res, err
	}
	return runners, res, nil
}

// CreateRunnersRegistrationToken Returns a token that you can pass to the config script. The token expires after one hour. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (s *ActionsService) CreateRunnersRegistrationToken(ctx context.Context, owner, repo string) (*ActionsToken, *github.Response, error) {
	return s.CreateRunnersTokenRequest(ctx, owner, repo, "registration")
}

// CreateRunnersRemoveToken Returns a token that you can pass to remove a self-hosted runner from a repository. The token expires after one hour. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (s *ActionsService) CreateRunnersRemoveToken(ctx context.Context, owner, repo string) (*ActionsToken, *github.Response, error) {
	return s.CreateRunnersTokenRequest(ctx, owner, repo, "remove")
}

// CreateRunnersTokenRequest Returns a create/remove token which can be used for the config script or remove script. The token expires after one hour. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (s *ActionsService) CreateRunnersTokenRequest(ctx context.Context, owner, repo, tokenType string) (*ActionsToken, *github.Response, error) {
	req, err := RestRequest(http.MethodPost, "%srepos/%s/%s/actions/runners/%s-token", s.client.BaseURL, owner, repo, tokenType)
	if err != nil {
		return nil, nil, err
	}

	var token ActionsToken
	res, err := s.client.Do(ctx, req, &token)
	if err != nil {
		return nil, res, err
	}
	return &token, res, nil
}

// DeleteRunners Forces the removal of a self-hosted runner from a repository. You can use this endpoint to completely remove the runner when the machine you were using no longer exists. Anyone with admin access to the repository can use this endpoint. GitHub Apps must have the administration permission to use this endpoint.
func (s *ActionsService) DeleteRunners(ctx context.Context, owner, repo string, runnerID int64) (*github.Response, error) {
	req, err := RestRequest(http.MethodDelete, "%srepos/%s/%s/actions/runners/%d", s.client.BaseURL, owner, repo, runnerID)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// RestRequest creates a http.Request for the given method and format string to build the url
func RestRequest(method string, format string, args ...interface{}) (*http.Request, error) {
	url := fmt.Sprintf(format, args...)
	return http.NewRequest(method, url, nil)
}
