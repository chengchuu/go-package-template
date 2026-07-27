# Go Package Template

A production-ready starting point for reusable Go packages. It includes an idiomatic public API,
tests and examples, source-derived API documentation, a responsive Bootstrap 5.3.8 website, an
examples page generated from real Go source, SEO metadata, installable PWA support, deterministic
validation, and GitHub Pages deployment.

The public package is at the module root. That keeps installation and imports concise while
`internal/site` prevents website tooling from becoming consumer API.

## Install

Go 1.25 or later is supported.

```bash
go get github.com/chengchuu/go-package-template
```

```go
import gopackage "github.com/chengchuu/go-package-template"
```

## Usage

```go
label, err := gopackage.FormatLabel(
	"go package",
	gopackage.WithStyle(gopackage.StyleTitle),
	gopackage.WithAffixes("[", "]"),
)
if err != nil {
	log.Fatal(err)
}
fmt.Println(label) // [Go Package]
```

Run the complete example with `go run ./examples/basic`.

## Website and documentation

- [Public website](https://chengchuu.github.io/go-package-template/)
- [Executable examples](https://chengchuu.github.io/go-package-template/examples/)
- [Generated Go API reference](https://chengchuu.github.io/go-package-template/api/)
- [Basic example](examples/basic/main.go)

The examples page renders `examples/basic/main.go` and `example_test.go` directly from the repository.
Run them with the Go toolchain to exercise the actual public package.

## Development

```bash
gofmt -w .
go mod tidy
go test ./...
go vet ./...
go build ./...
go run ./examples/basic
```

Generate and validate the final Pages artifact:

```bash
go run ./cmd/sitegen build
go run ./cmd/sitegen validate
```

Preview the production-like artifact at
`http://127.0.0.1:4173/go-package-template/` (including its scoped service worker):

```bash
go run ./cmd/sitegen preview
```

`dist/pages` is generated and ignored by Git. CI regenerates, validates, and uploads it; never edit
generated Pages files by hand.

## Reusing the template

Follow [CUSTOMIZE.md](CUSTOMIZE.md) in order. Project identity, routes, SEO, theme, manifest, and
cache settings are centralized in `internal/site/config.go`.

## License

The Go package and project-authored website code are available under the [MIT License](LICENSE).
Vendored browser-asset notices are in [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
