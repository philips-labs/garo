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
| Windows  | `choco install -y protoc`   |                                    |
| MacOSX   | `brew install protobuf`     |                                    |

To get an overview of the available make targets simply run the following:

```bash
$ make
clean-binaries       Cleans binaries
clean-diagrams       Cleans plantuml.jar and generated diagrams
clean                Clean binaries and diagrams
compile-agent        Compile garo-agent
compile-server       Compile garo-server
compile              Compile garo-agent and garo-server
coverage-html        Show coverage in browser
coverage-out         Show coverage in cli
diagrams             Generate diagrams in SVG and PNG format
download             Fetches go.mod dependencies via go mod download
install-protoc       Installs protoc
install-tools        Installs tools from tools.go
png-diagrams         Generate diagrams in PNG format
proto-gen            Generate protocol buffer implementations
svg-diagrams         Generate diagrams in SVG format
test-cover           Run tests and coverage
test                 Run tests
```

## Run

Github Testing implementation _(To Be Removed)_ can be run as following.

```bash
export GARO_GH_TOKEN=MYDUMMYPERSONALGHTOKEN
go build .
./garo my-gh-organization my-repository
```

### Server

The Server currently exposes a http server where agents can fetch a configuration. _The configuration specifics and remainder of features still has to be implemented._

```bash
$ bin/garo-server
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "GET", "route": "/"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "GET", "route": "/ping"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "OPTIONS", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "TRACE", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "POST", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "PUT", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "GET", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "CONNECT", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "HEAD", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "DELETE", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    DEBUG    server/api.go:75       Registering route  {"method": "PATCH", "route": "/twirp/philips.garo.garo.AgentConfigurationService/*"}
2020-02-27T14:45:03.921+0100    INFO     garo-server/root.go:63 Server is ready to handle requests  {"addr": ":8080"}
^C
2020-02-27T14:45:15.097+0100    INFO     server/server.go:81    Server is shutting down  {"reason": "interrupt"}
2020-02-27T14:45:15.097+0100    INFO     server/server.go:90    Server stopped
```

### Agent

The Agent currently connects to the server to fetch a configuration. _The configuration specifics and the remainder of features still has to be implemented._

```bash
bin/garo-agent
```

## Test

```bash
export GARO_GH_TOKEN=MYDUMMYPERSONALGHTOKEN
make test
```

To run the tests from _VSCode_ you will have to copy the `.env.example` and fill out your personal Github Token.

```shell
cp .env.example .env
```

To view the code coverage you can make use of `make coverage-out` or `make coverage-html`.

## Update diagrams

To update the diagrams you simply edit the plantuml files followed by running the according `make` target.

```bash
make clean-diagrams diagrams
```

> **NOTE:** due to `make` caching ensure to run the `clean-diagrams` target to regenerate existing `png` and `svg` files.

[cmake-3.16.2-win64-x64.msi]: https://github.com/Kitware/CMake/releases/download/v3.16.2/cmake-3.16.2-win64-x64.msi "Download cmake-3.16.2-win64-x64.msi"
[cmake-3.16.2-darwin-x86_64.dmg]: https://github.com/Kitware/CMake/releases/download/v3.16.2/cmake-3.16.2-Darwin-x86_64.dmg "Download cmake-3.16.2-Darwin-x86_64.dmg"
[graphviz-2.38.msi]: https://graphviz.gitlab.io/_pages/Download/windows/graphviz-2.38.msi "Download graphviz-2.38.msi"
[graphviz-2.42.2.tar.gz]: https://gitlab.com/graphviz/graphviz/-/archive/2.42.2/graphviz-2.42.2.tar.gz "Download graphviz-2.42.2.tar.gz"
