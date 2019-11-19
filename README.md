# Cloud Native Security Hub

![last commit](https://flat.badgen.net/github/last-commit/falcosecurity/cloud-native-security-hub-backend?icon=github) ![licence](https://flat.badgen.net/github/license/falcosecurity/cloud-native-security-hub-backend)

Cloud Native Security Hub is a platform for discovering and sharing rules and
configurations for cloud native security tools.

This repository contains the HTTP API and backend code who runs
https://securityhub.dev site

## Usage

This code requires a recent golang version (1.13) and it uses modules to handle
the dependencies.

### Configuration

You need to setup a couple of environment variables prior to run anything. As
long as SecurityHub uses plain YAML files to manage security resources, you
need to adjust its location:

* `RESOURCES_PATH`: Path to securityhub/resources directory
* `VENDOR_PATH`: Path to securityhub/vendors directory

### cmd/server

This is the HTTP API server and it will listen to requests on `8080` port.

```
$ go run cmd/server/main.go
```

## Contributing

Contributors are welcome! You will need a quick package overview to understand
some design decisions:

* `pkg/usecases`: You will find the entry points in the `pkg/usecases` directory.
  One action per file, modeled like a command.
* `pkg/resource` and `pkg/vendor`: This is the domain code for security resource
  and vendor. You will find the repositories, entities and value objects.
* `test`: All our code is test driven, in this directory we have some fixtures
  to avoid repeating test data in the test code.
* `web`: The web is just a delivery mechanism, is separated from the backend code,
  which can be used as a library if you need to. Is responsible to JSON
  marshalling and HTTP communications.
