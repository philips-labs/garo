package main

import (
	"context"
	"log"
	"os"

	"github.com/philips-labs/garo/filter"
	"github.com/philips-labs/garo/github"
)

func main() {
	token := os.Getenv("GARO_GH_TOKEN")
	if token == "" {
		log.Fatal("You need to export the environment variable GARO_GH_TOKEN")
	}
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("You need to provide a Github organization and repo as argument to GARO, e.g.:\n  %s my-org my-repo", os.Args[0])
	}

	org := args[0]
	repo := args[1]

	ctx := context.Background()
	oauthClient := github.PersonalAccessToken(ctx, token)
	client := github.NewClient(oauthClient)

	repos, _, err := client.Repositories.ListByOrg(ctx, org, nil)
	dieOnErr("Failed listing repos: %s", err)
	filtered := filter.Repositories(repos, filter.ByLanguage("go"))
	printJSON("Go Repositories: %s\n", filtered)

	factService := filter.Repositories(repos, filter.ByName(repo))[0]
	printJSON(repo+" Repository: %s\n", factService)

	runners, _, err := client.Actions.ListRunners(ctx, org, repo)
	dieOnErr("Failed listing runners: %s", err)
	printJSON("Runners: %s\n", runners)

	downloads, _, err := client.Actions.ListRunnersDownloads(ctx, org, repo)
	dieOnErr("Failed listing runner downloads: %s", err)
	printJSON("Runner downloads: %s\n", downloads)

	registrationToken, _, err := client.Actions.CreateRunnersRegistrationToken(ctx, org, repo)
	dieOnErr("Failed creating registration token: %s", err)
	printJSON("Create token: %s\n", registrationToken)

	removeToken, _, err := client.Actions.CreateRunnersRemoveToken(ctx, org, repo)
	dieOnErr("Failed creating remove token: %s", err)
	printJSON("Remove token: %s\n", removeToken)

	workflows, _, err := client.Actions.ListWorkflows(ctx, org, repo, nil)
	dieOnErr("Failed listing workflows: %s", err)
	printJSON("Workflows: %s\n", workflows)

	if workflows.TotalCount > 0 {
		workflow, _, err := client.Actions.GetWorkflowByID(ctx, org, repo, workflows.Workflows[0].ID)
		dieOnErr("Failed getting workflow by ID: %s", err)
		printJSON("Workflow: %s\n", workflow)
	}
}
