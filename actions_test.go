package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx    context.Context
	client *Client
)

func init() {
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		log.Fatal("You need to export the environment variable GH_TOKEN")
	}
	ctx = context.Background()
	oauthClient := PersonalAccessToken(ctx, token)
	client = NewClient(oauthClient)
}

func TestListActionRunners(t *testing.T) {
	assert := assert.New(t)
	org := "philips-internal"
	repo := "fact-service"
	runners, _, err := client.ListActionsRunners(ctx, org, repo)
	assert.NoError(err, "Failed listing action runners for %s/%s", org, repo)
	assert.NotEmpty(runners)
}

func TestListActionRunnerDownloads(t *testing.T) {
	assert := assert.New(t)
	org := "philips-internal"
	repo := "fact-service"
	downloads, _, err := client.ListActionsRunnersDownloads(ctx, org, repo)
	assert.NoError(err, "Failed listing action runner downloads for %s/%s", org, repo)
	assert.NotEmpty(downloads)
	assert.Len(downloads, 5, "Expected to have 5 downloads")
}
