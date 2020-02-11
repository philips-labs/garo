package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
)

func main() {
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		log.Fatal("You need to export the environment variable GH_TOKEN")
	}
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You need to provide a Github organization as argument to GARO")
	}

	ctx := context.Background()
	oauthClient := PersonalAccessToken(ctx, token)
	client := github.NewClient(oauthClient)

	repos, _, err := client.Repositories.ListByOrg(ctx, args[0], nil)
	dieOnErr("Failed listing repos: %s", err)
	printJSON("Repositories: %s\n", repos)
}
