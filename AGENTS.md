# AGENTS.md

Guidance for automated coding agents working in `go-package-template`.

## Scope and package contract

This repository is a reusable Go package template with a generated public website. The module path
is owned by `go.mod`; its mirrored build-time identity lives in `internal/site/config.go` and is
guarded by tests. The public package is the module-root `gopackage` package in `doc.go` and
`package.go`. Keep that root import path intentional and stable.

The package runtime uses only the Go standard library. Website code must not be imported by the
public Go package. Website changes must not silently alter the public Go package contract. A public
API change requires aligned GoDoc comments, behavior tests, example tests, `examples/basic`, README
usage, homepage usage, the generated examples page, and regenerated API documentation.

## Project structure

- `doc.go`, `package.go`: public package documentation and implementation.
- `package_test.go`, `example_test.go`: external-package behavior and executable documentation tests.
- `examples/basic`: runnable consumer-style example.
- `cmd/sitegen`: Go CLI for Pages build, validation, and production-like preview.
- `internal/site/config.go`: single editable website/project identity source; it is build tooling,
  never public API.
- `internal/site`: API extraction, metadata, SEO/PWA generation, icon generation, final-artifact
  validation, preview server, and deterministic tests.
- `site/templates`: crawlable homepage, examples, API, shared navigation, and layout sources.
- `site/assets/js`: shared theme, mobile navigation, and site/PWA registration behavior.
- `site/assets/css`: small Bootstrap-variable-based project styling.
- `site/assets/vendor`: website-only Bootstrap 5.3.8 assets.
- `site/assets/icons`: source favicon, PWA, maskable, Apple-touch, and Open Graph artwork copied
  unchanged into the generated site.
- `dist/pages`: canonical generated Pages artifact; ignored by Git and generated in CI.
- `.github/workflows/ci.yml`: package and final-site quality gates.
- `.github/workflows/pages.yml`: validated Pages artifact upload and deployment.
- `CUSTOMIZE.md`: ordered Go-specific post-copy conversion checklist.

## Stable website contract

The final artifact must retain `index.html`, `examples/index.html`, and `api/index.html`, deployed
below `/go-package-template/`. GitHub Pages paths must remain repository-subpath-aware. The API page
must come from actual exported root-package declarations; do not document `internal` packages as
consumer API.

All public pages share `navigation.tmpl`, `layout.tmpl`, Bootstrap initialization, `theme.js`, and
`site.js`. The examples page renders the checked-in runnable command and example test directly from
Go source; do not replace it with browser-side simulation or remote code execution. New public pages
must include crawlable HTML and accurate self-canonical metadata.

SEO metadata, JSON-LD, `robots.txt`, and `sitemap.xml` are generated from `SiteConfig`. PWA manifest,
icons, service-worker configuration, and precache routes are generated in Go. PWA caches and theme
storage keys must be project-specific. The service worker handles only same-origin GET requests in
the configured scope, uses network-first HTML, cache-first static assets, bounded runtime caching,
and prefix-limited cleanup.

## Commands and generated files

Run from the repository root:

```bash
gofmt -w .
go mod tidy
go test ./...
go vet ./...
go build ./...
go run ./examples/basic
go run ./cmd/sitegen build
go run ./cmd/sitegen validate
go run ./cmd/sitegen preview
```

`preview` rebuilds and validates before serving
`http://127.0.0.1:4173/go-package-template/`; use it for production-like PWA testing. Ordinary file
viewing does not register a worker.

Generated Pages output must not be edited by hand. Update Go configuration/generators, site source,
or package GoDoc, then rebuild. CI generates `dist/pages`; do not commit it. If validation dirties
generated output, leave source correct and regenerate rather than patching output.

## Change discipline

Keep package examples, README, tests, and API documentation aligned. Add deterministic tests for
new derivation, metadata, route, scope, or validation behavior. Keep checks independent of external
network services. Preserve accessible focus, semantic landmarks, initial content, one `h1` per page,
and unique metadata. Follow `CUSTOMIZE.md` when changing identity so module, repository, Pages,
theme, PWA, icon, workflow, and documentation values move together.

Do not add runtime dependencies without a documented consumer benefit. Do not add Node.js merely
for site generation; Bootstrap is vendored website input. Do not publish releases or tags unless
explicitly requested.
