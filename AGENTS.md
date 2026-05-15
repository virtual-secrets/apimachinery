# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.virtual-secrets.dev/apimachinery`: the Virtual Secrets project's Kubernetes API types, generated clients, OpenAPI definitions, and CRD manifests. This is a library repo — it produces no binary. Other Virtual Secrets components (the operator/server, CSI provider, etc.) depend on these types.

## Architecture

- `apis/` — Kubernetes API type definitions, split by group:
  - `apis/config/v1alpha1/` — group `config.virtual-secrets.dev` (e.g. `SecretStore`, `SecretMetadata`).
  - `apis/virtual/v1alpha1/` — group `virtual-secrets.dev` (e.g. `Secret`, `SecretMount`).
  - Each group has `install/` (scheme registration) and `fuzzer/` (round-trip fuzz helpers).
  - `groupversion_info.go`, `*_types.go` are hand-written; `zz_generated.*.go` and `openapi_generated.go` are generated — do not edit by hand.
- `client/` — generated typed clientsets, listers, informers (k8s.io/client-go style). Generated; do not edit.
- `crds/` — CRD YAML manifests generated from the API types, plus `lib.go` exposing them via go embed.
- `openapi/` — generated OpenAPI spec for the API server.
- `hack/` — codegen and build helpers:
  - `hack/gencrd/` — Go program that drives CRD generation.
  - `hack/scripts/`, `hack/build.sh`, `hack/fmt.sh`, `hack/test.sh` — invoked by the Makefile inside Docker.
- `vendor/` — vendored dependencies; `go mod tidy && go mod vendor` keeps it in sync.

API group/version pairs are defined in the Makefile as `API_GROUPS := config:v1alpha1 virtual:v1alpha1`; codegen targets fan out over this list.

## Common commands

All build/test/lint targets run inside the `ghcr.io/appscode/golang-dev` Docker image — you generally do not need a local Go toolchain, but Docker must be running.

- `make ci` — full pipeline used by GitHub Actions: `verify check-license lint build unit-tests`. Run this before opening a PR.
- `make gen` — regenerate everything: `clientset manifests openapi`. Run after any change to `apis/**/*_types.go` or `apis/**/groupversion_info.go`.
- `make manifests` — regenerate CRDs only (`gen-crds patch-crds label-crds`).
- `make clientset` — regenerate `client/` only.
- `make openapi` — regenerate OpenAPI definitions only.
- `make fmt` — gofmt + goimports across `apis client crds hack/gencrd`.
- `make lint` — golangci-lint.
- `make unit-tests` (or `make test`) — run Go unit tests.
- `make build` — build a binary (mostly a no-op sanity check for this lib repo).
- `make verify` — `verify-gen` (re-run `gen fmt` and confirm tree is clean) + `verify-modules` (`go mod tidy && go mod vendor` clean).
- `make add-license` — apply Apache-2.0 license headers to source files.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/virtual/v1alpha1/... -run TestName -v
```

CI also runs CRD apply tests against Kind clusters at Kubernetes versions `v1.26.15`, `v1.28.15`, `v1.30.10`, `v1.32.2` (`kubectl create -R -f ./crds`). If you add or change a CRD, ensure it still applies cleanly on these versions.

## Conventions

- Module path is `go.virtual-secrets.dev/apimachinery` (vanity URL); imports must use that, not the GitHub URL.
- Do not hand-edit any file starting with `zz_generated.` or anything under `client/`, `openapi/openapi_generated.go`, or `crds/*.yaml`. Change the source in `apis/**/*_types.go` and re-run `make gen`.
- Kubebuilder/CRD options: `crd:maxDescLen=0,generateEmbeddedObjectMeta=true,allowDangerousTypes=true` (see `CRD_OPTIONS` in Makefile) — keep this in mind when adding new types.
- All contributions must be signed off (`git commit -s`) — the repo carries a DCO file.
- License: Apache-2.0. New files need the standard AppsCode header (`make add-license` will add it).
- Dependabot (`.github/dependabot.yml`) manages dependency PRs; the `release-tracker.yml` and `update-crds.yml` workflows automate release bookkeeping.
