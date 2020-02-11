# Github actions runners orchestrator

GitHub allows developers to run GitHub Actions workflows on your own runners. This tool allows you to deploy self hosted runners for your repositories.

## Important notes

> GitHub [recommends](https://help.github.com/en/github/automating-your-workflow-with-github-actions/about-self-hosted-runners#self-hosted-runner-security-with-public-repositories) that you do **NOT** use self-hosted runners with public repositories, for security reasons.

## Run

GARO currently supports listing a page of organization repositories.

```bash
export GH_TOKEN=MYDUMMYPERSONALGHTOKEN
./garo my-gh-organization
```

## Test

```bash
export GH_TOKEN=MYDUMMYPERSONALGHTOKEN
go test -v ./...
```
