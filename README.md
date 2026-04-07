# pulumi-go-playground

A benchmarking ground for the **code-generated Go SDKs** that Pulumi ships for each of its providers.

Pulumi's Go SDKs are generated from provider schemas, which can produce very large packages (thousands of resource types per cloud). The goal of this repo is to measure how much each generated SDK actually costs to **download, compile, and `pulumi preview`** in isolation, and to track how those numbers move over time as providers are upgraded.

## Quick start

The repo uses [mise](https://mise.jdx.dev/) to pin tool versions (`go`, `pulumi`) and to expose per-provider tasks. Once you have mise installed:

```bash
mise install                       # provision Go + Pulumi at the pinned versions
mise run -C <provider> build       # build the provider's bench binary (-> <provider>/dist/infra)
```

For example, to build the cloudflare bench locally:

```bash
mise run -C cloudflare build
```

Each `<provider>/.mise.toml` defines a `build` task that runs `go build -v -o dist/infra .` — the same command CI runs in its `build` job, alongside `pulumi preview` against the local file backend. See [What we're measuring](#what-were-measuring).

## How it's organised

- **`PROVIDERS.md`** — inventory of every `github.com/pulumi/*` provider that ships a Go SDK, with the latest released version and a ready-to-paste `go get` line. Source of truth for what's bench-able.
- **`<provider>/`** — one folder per provider being benchmarked. Each folder is its own self-contained Go module with a minimal `main.go` that does nothing more than initialise the provider via `NewProvider`. This isolates the SDK's cost from any user code.
- **`.github/workflows/<provider>-preview.yml`** — per-provider GitHub Actions workflow that runs `pulumi preview` against the local file backend (`cloud-url: file://~`, no Pulumi Cloud token needed). Workflow timings are the headline benchmark signal.
- **`.claude/skills/bootstrap-provider/`** — Claude Code skill that scaffolds a new provider folder + workflow in one step. See below.

## Adding a provider to the benchmark

From inside Claude Code:

```
/bootstrap-provider <name>
```

For example, `/bootstrap-provider aws`. The skill:

1. Looks up the module path and pinned version from `PROVIDERS.md`.
2. Creates `<name>/` as a fresh Go module with the provider + `pulumi/sdk/v3` as dependencies.
3. Writes a minimal `main.go` calling `<pkg>.NewProvider`, plus a `Pulumi.yaml`.
4. Verifies the scaffold compiles (`go mod tidy && go build ./...`).
5. Writes `.github/workflows/<name>-preview.yml` so CI starts benchmarking it on the next push.

If the provider doesn't expose a `NewProvider` constructor, the skill stops without creating anything.

## What we're measuring

The per-provider workflows give us, per SDK version:

- **`go mod download`** — wall-clock cost of fetching the generated SDK (proxy for SDK size on disk).
- **`go build`** — compile time for a program that touches only `NewProvider`. This is the smallest possible "blast radius" of importing the SDK.
- **`pulumi preview`** — end-to-end cost of running the Pulumi engine against the trivial program, including provider plugin download.

Comparing these numbers across providers (and across versions of the same provider) is the point of the playground.
