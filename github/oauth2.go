package github

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// PersonalAccessToken creates a new Oauth2 http client utilizing the personal access token
func PersonalAccessToken(ctx context.Context, token string) *http.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	return oauth2.NewClient(ctx, tokenSource)
}
