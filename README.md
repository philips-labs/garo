# Github actions runners orchestrator

![CI](https://github.com/philips-labs/garo/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/philips-labs/garo/branch/develop/graph/badge.svg)](https://codecov.io/gh/philips-labs/garo)

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

To run the tests from *VSCode* you will have to copy the `.env.example` and fill out your personal Github Token.

```shell
cp .env.example .env
```
