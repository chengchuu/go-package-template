# Customize this Go package template

Work through this checklist in order. Change source and configuration first, then regenerate
`dist/pages`; generated Pages files are never authoritative and must not be edited by hand.

## 1. Replace the Go module path

Change the `module` directive in `go.mod`, then update `modulePath` in `site.config.json`, Go imports,
examples, and README commands. Use the final source-control import path, including a major-version
suffix when required by Go modules.

## 2. Rename the package

Rename the root `package gopackage` declarations and import alias. This template intentionally uses
an idiomatic root package; move it to `pkg/<name>` only when that extra import-path segment is a
deliberate consumer contract. Update `packageName` in `site.config.json` either way.

## 3. Replace the exported sample API

Replace `Formatter`, `Labeler`, `FormatLabel`, options, styles, and sentinel errors with the real
library surface. Keep exported GoDoc comments, focused behavior tests, examples, and error contracts
aligned. Avoid exposing website configuration from the package.

## 4. Set the repository owner and name

Edit `repository.owner` and `repository.name` in `site.config.json`. Repository URL, Pages path,
production URL, theme storage key, and service-worker cache prefix are derived and validated by Go.

## 5. Confirm the Pages base path and production URL

`PagesBasePath` and `PagesURL` are derived from the repository identity. For a project repository,
keep the `/repository-name/` path. If the deployment is intentionally a user site or custom domain,
update the derivation functions, tests, canonical expectations, and workflow together.

## 6. Replace the package description

Update the shared top-level `description` in `site.config.json`, root GoDoc, and the README
introduction. Each generated page reuses this description while retaining its own configured title.
Keep claims limited to tested capabilities.

## 7. Rewrite README examples

Update installation, import, usage, supported Go version, example path, and API links. Compile every
snippet through an example test or runnable command.

## 8. Rewrite website content

Edit `site/templates/home.tmpl`, `examples.tmpl`, and supporting CSS or JavaScript. Preserve stable
routes unless an intentional migration is planned. Every public page needs useful initial HTML,
one descriptive `h1`, and accessible controls.

## 9. Change the theme storage key

The storage key is derived from `repository.name`. Review `theme.colorLight`, `theme.colorDark`, and
the coordinated light and dark primary palettes in `site.config.json` together. Keep the supported
preferences exactly `system`, `light`, and `dark`, and keep browser `theme-color` synchronized with
the resolved mode.

## 10. Update manifest identity

Review `displayName`, `pwa.shortName`, description, and colors in `site.config.json`. Generated
`start_url` and `scope` must remain equal to the derived Pages base path for project-site deployment.

## 11. Change the service-worker cache prefix

The cache prefix is derived from `repository.name`. Review precached routes after adding or removing
public pages; cleanup must continue to target only caches with this project-specific prefix.

## 12. Replace favicon and PWA icons

Replace the 32px favicon, 192px and 512px application icons, padded 512px maskable icon, and 1200×630
Open Graph image under `site/assets/icons`. Keep the filenames in `site.config.json`, actual
dimensions, manifest declarations, maskable safe zones, Apple-touch link, and social metadata aligned.

## 13. Review GitHub Actions settings

Update default branches, Go version, Pages environment settings, and action versions in
`.github/workflows/ci.yml` and `pages.yml`. Keep package checks before artifact upload and retain only
the permissions required for Pages. Keep `goVersion` in `site.config.json` aligned with the `go`
directive in `go.mod`; the site generator rejects version drift.

## 14. Replace license and copyright

Choose the intended license, replace `LICENSE`, and update `license` in `site.config.json`, README,
footer, and copyright owner. Preserve third-party notices for retained Bootstrap files.

## 15. Regenerate API documentation

The API page comes from the exported root package via Go AST and GoDoc data. After public API edits,
run the site build and inspect `dist/pages/api/index.html`; do not hand-author declarations in the
template.

## 16. Find remaining template identity

Search the source tree for `go-package-template`, `gopackage`, `chengchuu`, old URLs, icon names,
theme keys, and cache prefixes. Decide which historical references are intentional. The artifact
validator also rejects known unresolved template markers and stale reference-project identity.

## Final verification

```bash
gofmt -w .
go mod tidy
go test ./...
go vet ./...
go build ./...
go list ./...
go list ./... | xargs -n 1 go doc
go run ./examples/basic
go run ./cmd/sitegen build
go run ./cmd/sitegen validate
```

`go doc` does not expand the `./...` package pattern itself, so the `go list` pipeline runs it once
for each package in the module.

Then run `go run ./cmd/sitegen preview` and inspect the home, examples, and API routes below the
configured project path. Confirm direct refresh, offline fallback, install metadata, theme changes,
mobile navigation, canonical URLs, and internal links before enabling Pages for the copied repository.
