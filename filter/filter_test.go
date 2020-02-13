package filter_test

import (
	"context"
	"log"
	"os"
	"testing"

	gh "github.com/google/go-github/v29/github"
	"github.com/stretchr/testify/assert"

	"github.com/philips-labs/garo/filter"
	"github.com/philips-labs/garo/github"
)

var (
	ctx    context.Context
	client *github.Client
)

func init() {
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		log.Fatal("You need to export the environment variable GH_TOKEN")
	}
	ctx = context.Background()
	oauthClient := github.PersonalAccessToken(ctx, token)
	client = github.NewClient(oauthClient)
}

func TestFilterRepositories(t *testing.T) {
	assert := assert.New(t)
	org := "philips-labs"
	repos, _, err := client.Repositories.ListByOrg(ctx, org, &gh.RepositoryListByOrgOptions{Type: "public"})
	assert.NoError(err, "Failed listing repositories for %s", org)

	filtered := filter.Repositories(repos, filter.ByLanguage("go"))
	assert.NotEmpty(filtered, "Expected to have a bunch of public repositories with language 'Go'.")
	assert.NotEqual(len(repos), len(filtered), "Expected to have less repositories after applying language filter")

	filtered = filter.Repositories(repos, filter.ByName("garo"))
	assert.Len(filtered, 1)
}
