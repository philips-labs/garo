package main

import (
	"context"
	"log"
	"os"

	"github.com/philips-labs/garo/github"
)

func main() {
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		log.Fatal("You need to export the environment variable GH_TOKEN")
	}
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("You need to provide a Github organization and repo as argument to GARO, e.g.:\n  %s my-org my-repo", os.Args[0])
	}

	ctx := context.Background()
	oauthClient := github.PersonalAccessToken(ctx, token)
	client := github.NewClient(oauthClient)

	repos, _, err := client.Repositories.ListByOrg(ctx, args[0], nil)
	dieOnErr("Failed listing repos: %s", err)
	printJSON("Repositories: %s\n", repos)

	runners, _, err := client.Actions.ListRunners(ctx, args[0], args[1])
	dieOnErr("Failed listing runners: %s", err)
	printJSON("Runners: %s\n", runners)

	downloads, _, err := client.Actions.ListRunnersDownloads(ctx, args[0], args[1])
	dieOnErr("Failed listing runner downloads: %s", err)
	printJSON("Runner downloads: %s\n", downloads)

	registrationToken, _, err := client.Actions.CreateRunnersRegistrationToken(ctx, args[0], args[1])
	dieOnErr("Failed creating registration token: %s", err)
	printJSON("Create token: %s\n", registrationToken)

	removeToken, _, err := client.Actions.CreateRunnersRemoveToken(ctx, args[0], args[1])
	dieOnErr("Failed creating remove token: %s", err)
	printJSON("Remove token: %s\n", removeToken)
}
