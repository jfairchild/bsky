# Copilot Instructions

## Project Overview

This is a Go application that posts messages to [Bluesky](https://bsky.social) (AT Protocol) social network. It uses the [`indigo`](https://github.com/bluesky-social/indigo) library for Bluesky API interactions and OpenTelemetry for observability.

## Repository Structure

- `main.go` — Entry point; contains `PostToBsky` and `PrintISODateTime` functions
- `otel.go` — OpenTelemetry SDK setup (`setupOTelSDK`)
- `main_test.go` — Unit tests
- `go.mod` / `go.sum` — Go module dependencies
- `.mise.toml` — [Mise](https://mise.jdx.dev/) tool configuration (Go version, env vars)
- `.github/workflows/ci.yml` — CI pipeline (runs `go test ./...` on pull requests)

## Build, Test, and Lint

```bash
# Run all tests
go test ./...

# Build the binary
go build ./...

# Format code (always run before committing)
gofmt -w .

# Vet code for common issues
go vet ./...
```

## Coding Conventions

- Follow standard Go conventions and idioms (effective Go, Go code review comments)
- Use `fmt.Errorf("...: %w", err)` for error wrapping
- All functions that can fail should return an `error` as the last return value
- Prefer `context.Context` as the first argument for functions that perform I/O or long-running operations
- Use the `go-logr/logr` interface for structured logging; obtain the logger from context via `logr.FromContextOrDiscard(ctx)`
- Run `gofmt` before committing; code must be properly formatted
- Avoid adding new dependencies unless strictly necessary; prefer the standard library or already-imported packages

## Environment Variables

The following environment variables are used by OpenTelemetry (configured via `.mise.toml`):

| Variable | Description |
|---|---|
| `OTEL_SERVICE_NAME` | Service name reported to the telemetry backend |
| `OTEL_TRACES_SAMPLER` | Sampling strategy (`always_on`, `always_off`, `traceidratio`) |

## Dependencies

Key libraries used in this project:

- [`github.com/bluesky-social/indigo`](https://github.com/bluesky-social/indigo) — Bluesky/AT Protocol client
- [`go.opentelemetry.io/otel`](https://opentelemetry.io/docs/languages/go/) — OpenTelemetry SDK
- [`github.com/go-logr/logr`](https://github.com/go-logr/logr) — Structured logging interface
- [`github.com/urfave/cli/v2`](https://github.com/urfave/cli) — CLI framework

## CI

The CI pipeline (`ci.yml`) runs on every pull request to `main`. It:

1. Checks out the code
2. Sets up Go using the version from `go.mod`
3. Runs all tests with `go test ./...`

Ensure all tests pass locally before opening a pull request.
