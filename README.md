# Github actions runners orchestrator

![CI](https://github.com/philips-labs/garo/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/philips-labs/garo/branch/develop/graph/badge.svg)](https://codecov.io/gh/philips-labs/garo)

GitHub allows developers to run GitHub Actions workflows on your own runners. This tool allows you to deploy self hosted runners for your repositories.

[Documentation](docs/README.md)

## Important notes

> GitHub [recommends](https://help.github.com/en/github/automating-your-workflow-with-github-actions/about-self-hosted-runners#self-hosted-runner-security-with-public-repositories) that you do **NOT** use self-hosted runners with public repositories, for security reasons.

## Prerequisites

This project makes use of CMake to automate a few tasks in development. Therefore it is recommended to install **CMake**. Furthermore the Cmake tasks depend on Graphviz to generate the diagrams.

| platform | install                     | url                                |
| -------- | --------------------------- | ---------------------------------- |
| Windows  | `choco install -y cmake`    | [cmake-3.16.2-win64-x64.msi][]     |
| MacOSX   | `brew install cmake`        | [cmake-3.16.2-Darwin-x86_64.dmg][] |
| Windows  | `choco install -y graphviz` | [graphviz-2.38.msi][]              |
| MacOSX   | `brew install graphviz`     | [graphviz-2.42.2.tar.gz][]         |

To get an overview of the available make targets simply run the following:

```bash
$ make
clean                Cleans plantuml.jar and generated diagrams
diagrams             Generate diagrams in SVG and PNG format
png-diagrams         Generate diagrams in PNG format
svg-diagrams         Generate diagrams in SVG format
```

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

To run the tests from _VSCode_ you will have to copy the `.env.example` and fill out your personal Github Token.

```shell
cp .env.example .env
```

## Update diagrams

To update the diagrams you simply edit the plantuml files followed by running the according `make` target.

```bash
make clean diagrams
```

> **NOTE:** due to `make` caching ensure to run the `clean` target to regenerate existing `png` and `svg` files.

[cmake-3.16.2-win64-x64.msi]: https://github.com/Kitware/CMake/releases/download/v3.16.2/cmake-3.16.2-win64-x64.msi "Download cmake-3.16.2-win64-x64.msi"
[cmake-3.16.2-darwin-x86_64.dmg]: https://github.com/Kitware/CMake/releases/download/v3.16.2/cmake-3.16.2-Darwin-x86_64.dmg "Download cmake-3.16.2-Darwin-x86_64.dmg"
[graphviz-2.38.msi]: https://graphviz.gitlab.io/_pages/Download/windows/graphviz-2.38.msi "Download graphviz-2.38.msi"
[graphviz-2.42.2.tar.gz]: https://gitlab.com/graphviz/graphviz/-/archive/2.42.2/graphviz-2.42.2.tar.gz "Download graphviz-2.42.2.tar.gz"
