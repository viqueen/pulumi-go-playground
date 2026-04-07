---
name: bootstrap-provider
description: Bootstrap a new Go folder that initialises a Pulumi provider, ready for `pulumi preview`, plus a token-free GitHub Actions workflow that runs the preview on every PR. Use when the user says "bootstrap <provider>", "scaffold a pulumi go project for <provider>", or similar. Takes a provider short name (e.g. "aws", "gcp", "kubernetes") as argument.
---

# bootstrap-provider

Scaffolds a self-contained Go module under `./<provider>/` that initialises the requested Pulumi provider, plus a per-provider GitHub Actions workflow that runs `pulumi preview` against the local file backend.

## Inputs

- `$ARGUMENTS` — the provider short name as it appears in `PROVIDERS.md` (e.g. `aws`, `gcp`, `cloudflare`, `kubernetes`). This is the package name, not the repo name. The repo name is `pulumi-<provider>`.

If the argument is missing, ask the user which provider to bootstrap and stop.

## Steps

1. **Resolve the provider** by reading `PROVIDERS.md` at the repo root.
   - Find the row for `pulumi-<provider>`.
   - Extract the `go get` line — it gives you the module path (with `/v<major>` if present) and the version tag.
   - If the provider is not in `PROVIDERS.md` or is listed under "No Go SDK", stop and tell the user.
   - If the provider has no tagged release (e.g. `pulumi-hyperv`, `pulumi-ucloud`), stop and tell the user.

2. **Verify the provider exposes `NewProvider`.** Inspect the upstream Go SDK (e.g. `gh api repos/pulumi/pulumi-<provider>/contents/sdk[/v<major>]/go/<pkg>` and look for a `provider.go` that defines `NewProvider`). If it does not, **stop and tell the user — do not bootstrap anything.**

3. **Refuse to overwrite.** If `./<provider>/` already exists, stop and ask the user before doing anything.

4. **Create the folder and Go module.** Run from the repo root:
   ```bash
   mkdir <provider>
   cd <provider>
   go mod init github.com/viqueen/pulumi-go-playground/<provider>
   ```

5. **Add dependencies.** Use the exact `go get` line from `PROVIDERS.md` for the provider, plus the Pulumi Go SDK:
   ```bash
   go get github.com/pulumi/pulumi/sdk/v3@latest
   go get <module-path-from-PROVIDERS.md>@<version>
   ```

6. **Write `main.go`** using the Write tool. Template:
   ```go
   package main

   import (
       "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
       "<provider-import-path>"
   )

   func main() {
       pulumi.Run(func(ctx *pulumi.Context) error {
           _, err := <pkg>.NewProvider(ctx, "<provider>", &<pkg>.ProviderArgs{})
           if err != nil {
               return err
           }
           return nil
       })
   }
   ```

7. **Write a minimal `Pulumi.yaml`** in `<provider>/`:
   ```yaml
   name: <provider>-bootstrap
   runtime: go
   description: Bootstrap project for the pulumi-<provider> Go SDK
   ```

8. **Tidy and build** to prove the scaffold compiles:
   ```bash
   go mod tidy
   go build ./...
   ```
   Don't leave a broken scaffold. Fix build errors before moving on.

9. **Write the GitHub Actions workflow** to `.github/workflows/<provider>-preview.yml`:
   ```yaml
   name: <provider> preview

   on:
     pull_request:
       paths:
         - "<provider>/**"
         - ".github/workflows/<provider>-preview.yml"
     push:
       branches: [main]
       paths:
         - "<provider>/**"
         - ".github/workflows/<provider>-preview.yml"

   jobs:
     preview:
       runs-on: ubuntu-24.04
       steps:
         - uses: actions/checkout@v4
         - uses: actions/setup-go@v5
           with:
             go-version-file: <provider>/go.mod
         - uses: pulumi/actions@v6
           env:
             PULUMI_CONFIG_PASSPHRASE: ${{ secrets.PULUMI_CONFIG_PASSPHRASE }}
           with:
             command: preview
             stack-name: dev
             work-dir: <provider>
             cloud-url: file://~
             upsert: true
   ```
   - Scope `paths` to the provider folder.

10. **Report** to the user: the folder created, the version pinned, the workflow path, and that the next CI run will preview it.

## Notes

- For `pulumi-azure-native` use the `pulumi-azure-native-sdk` repo and pick a single service module (ask the user which one if not specified).
- Keep the scaffold minimal — no extra resources, no config wiring, no README.
