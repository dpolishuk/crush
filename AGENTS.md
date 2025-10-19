# Repository Guidelines

## Project Structure & Module Organization
- Root contains `main.go`, `go.mod`, `Taskfile.yaml`, and user-facing docs like `README.md` and `CRUSH.md`.
- Application code lives under `internal/`, split by domain (for example `internal/app`, `internal/commandhistory`, `internal/tui`).
- Database queries and migrations are in `internal/db/` and `internal/db/migrations/`.
- Tests accompany source files as `*_test.go`; integration-style setup (e.g., DB-backed) resides in the same package directory.

## Build, Test, and Development Commands
- `task build` (or `go build .`) compiles the TUI binary into `./crush`.
- `task run -- "prompt"` launches the app with live changes; use `task dev` to enable profiling instrumentation.
- `task test` (alias for `go test ./...`) runs the full Go test suite.
- `task fmt` applies `gofumpt` formatting; `task lint` runs `golangci-lint` with the repo config.

## Coding Style & Naming Conventions
- Go files must be formatted with `gofumpt -w .`; CI expects gofumpt + gofmt output.
- Follow Go naming: exported identifiers use `CamelCase`, private helpers use `camelCase`, and tests use `TestXxx`.
- Keep packages focused; prefer short files with clear responsibilities inside the relevant `internal/<domain>` directory.
- Commit generated SQL via `sqlc`; do not hand-edit `internal/db/*.sql.go`.

## Testing Guidelines
- Use the standard Go testing package; table-driven tests are preferred.
- Name tests `TestFeatureBehavior`; integration helpers may live in `_test.go` files alongside code.
- For DB-dependent tests, leverage temporary directories and `db.Connect` (see `internal/commandhistory/service_test.go`).
- Run `go test ./...` before submitting; add focused commands like `go test ./internal/tui/components/chat/editor`.

## Commit & Pull Request Guidelines
- Follow Conventional Commits (`feat:`, `fix(tui):`, `chore:`); scopes mirror directory names when helpful.
- Each PR should describe behavior changes, note DB schema updates, and mention new CLI flags or config.
- Link related issues and include screenshots/GIFs for UI or TUI changes.
- Rebase on `main`, ensure translations/migrations are documented, and request review once lint and tests pass locally.
