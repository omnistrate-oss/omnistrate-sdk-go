# Go API client for Omnistrate

![CI](https://github.com/omnistrate-oss/omnistrate-sdk-go/actions/workflows/build.yml/badge.svg)

## Overview

[omnistrate-sdk-go](https://github.com/omnistrate-oss/omnistrate-sdk-go) is an auto-generated Go API client for interacting with [Omnistrate](https://omnistrate.com/) APIs. It is generated using the [OpenAPI Generator](https://openapi-generator.tech) from Omnistrate's OpenAPI specifications.

## Packages

The SDK is split into two packages, each targeting a different set of Omnistrate APIs:

| Package | Import Path | Description |
|---------|-------------|-------------|
| [v1](/v1/README.md) | `github.com/omnistrate-oss/omnistrate-sdk-go/v1` | Service Registration, Build, and Consumption APIs |
| [fleet](/fleet/README.md) | `github.com/omnistrate-oss/omnistrate-sdk-go/fleet` | Fleet management APIs (host clusters, operations, inventory, webhooks, etc.) |

Each package has its own detailed documentation with a full list of API endpoints and models.

## Installation

```sh
go get github.com/omnistrate-oss/omnistrate-sdk-go@latest
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	v1 "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	cfg := v1.NewConfiguration()
	client := v1.NewAPIClient(cfg)

	// Authenticate
	ctx := context.WithValue(context.Background(), v1.ContextAccessToken, "your-token")

	// Use the client
	providers, _, err := client.CloudProviderApiAPI.CloudProviderApiListCloudProvider(ctx).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(providers)
}
```

## How It Works

### Code Generation

The entire SDK is auto-generated from Omnistrate's OpenAPI specs. The generation pipeline is driven by the `Makefile`:

1. **OpenAPI specs** are fetched from `https://api.omnistrate.cloud/2022-09-01-00/openapi.yaml` (v1) and `.../fleet/openapi.yaml` (fleet).
2. **openapi-generator** produces Go client code into `v1/` and `fleet/` using custom Mustache templates from [`templates/go/`](/templates/go/).
3. The generated code includes typed API clients, request/response models, configuration, and documentation.

### Key Make Targets

| Target | Description |
|--------|-------------|
| `make all` | Clean, regenerate, tidy, and build |
| `make build` | Build the Go SDK |
| `make clean` | Remove generated `v1/` and `fleet/` directories |
| `make gen-go-sdk` | Generate the v1 API client |
| `make gen-fleet-go-sdk` | Generate the fleet API client |
| `make validate` | Validate OpenAPI specs |
| `make tidy` | Run `go mod tidy` |
| `make download` | Download Go module dependencies |

### CI/CD

The [build workflow](/.github/workflows/build.yml) runs on pushes and PRs to `main`:
- Checks out the code, sets up Go, downloads dependencies, and runs `make build`.
- On push to `main`, automatically bumps the version tag and publishes a new release.

## Project Structure

```
├── v1/                  # Generated v1 API client (Service Registration & Consumption)
│   ├── api_*.go         # API endpoint implementations
│   ├── model_*.go       # Request/response model structs
│   ├── client.go        # APIClient struct
│   ├── configuration.go # Configuration (server URL, auth, headers)
│   └── docs/            # Generated API & model documentation
├── fleet/               # Generated Fleet API client
│   ├── api_*.go         # Fleet API endpoint implementations
│   ├── model_*.go       # Fleet model structs
│   ├── client.go        # Fleet APIClient struct
│   ├── configuration.go # Fleet configuration
│   └── docs/            # Generated Fleet documentation
├── templates/go/        # Custom Mustache templates for code generation
├── src/                 # OpenAPI Generator meta scaffolding (unused)
├── Makefile             # Build, generate, and validate targets
└── .github/workflows/   # CI/CD (build + auto-tag)
```

## Contributing

Want to contribute? Awesome! You can find information about contributing to this
project in the [CONTRIBUTING](/CONTRIBUTING.md) page.

## About Omnistrate

[Omnistrate](https://omnistrate.com/) is the operating system for your SaaS,
offering enterprise-grade capabilities: automated provisioning, serverless
capabilities, auto-scaling, billing, monitoring, centralized logging,
self-healing, intelligent patching and much more!
