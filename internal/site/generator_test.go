package site

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestBuildCreatesAndValidatesStableRoutes(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	cfg := testConfig(t)
	cfg.OutputDir = "dist/test-pages"
	t.Cleanup(func() { _ = os.RemoveAll(filepath.Join(root, cfg.OutputDir)) })

	if _, err := Build(root, cfg); err != nil {
		t.Fatal(err)
	}
	for _, route := range []string{"index.html", "examples/index.html", "api/index.html"} {
		if _, err := os.Stat(filepath.Join(root, cfg.OutputDir, filepath.FromSlash(route))); err != nil {
			t.Errorf("route %s: %v", route, err)
		}
	}
	if _, err := os.Stat(filepath.Join(root, cfg.OutputDir, "playground", "index.html")); !os.IsNotExist(err) {
		t.Errorf("obsolete playground route still exists: %v", err)
	}
	examplesHTML, err := os.ReadFile(filepath.Join(root, cfg.OutputDir, "examples", "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	for _, sourceDeclaration := range []string{"func main()", "func ExampleFormatLabel()"} {
		if !strings.Contains(string(examplesHTML), sourceDeclaration) {
			t.Errorf("generated examples page is missing %q", sourceDeclaration)
		}
	}
	if _, err := ValidateArtifact(root, cfg); err != nil {
		t.Fatal(err)
	}
	homePath := filepath.Join(root, cfg.OutputDir, "index.html")
	homeHTML, err := os.ReadFile(homePath)
	if err != nil {
		t.Fatal(err)
	}
	for _, themeContract := range []string{
		`--project-primary-contrast: #141414;`,
		`--project-primary-state-contrast: #ffffff;`,
		`--bs-link-color: var(--project-primary-hover);`,
		`--bs-link-hover-color: var(--project-primary-active);`,
	} {
		if !strings.Contains(string(homeHTML), themeContract) {
			t.Errorf("generated homepage is missing accessible theme contract %q", themeContract)
		}
	}
	if !strings.Contains(string(homeHTML), `class="d-flex flex-column align-items-start gap-2 mt-4"`) {
		t.Error("generated homepage does not keep the source link below the primary hero actions")
	}
	expectedImport := `import ` + cfg.PackageName + ` "` + cfg.ModulePath + `"`
	staleImport := strings.Replace(string(homeHTML), expectedImport, `import stale "`+cfg.ModulePath+`"`, 1)
	if staleImport == string(homeHTML) {
		t.Fatalf("generated homepage is missing %q", expectedImport)
	}
	if err := os.WriteFile(homePath, []byte(staleImport), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ValidateArtifact(root, cfg); err == nil || !strings.Contains(err.Error(), "package import example does not match configuration") {
		t.Fatalf("ValidateArtifact error = %v, want stale package import failure", err)
	}
	if err := os.WriteFile(homePath, homeHTML, 0o644); err != nil {
		t.Fatal(err)
	}
	apiPath := filepath.Join(root, cfg.OutputDir, "api", "index.html")
	apiHTML, err := os.ReadFile(apiPath)
	if err != nil {
		t.Fatal(err)
	}
	staleImport = strings.Replace(string(apiHTML), expectedImport, `import stale "`+cfg.ModulePath+`"`, 1)
	if staleImport == string(apiHTML) {
		t.Fatalf("generated API page is missing %q", expectedImport)
	}
	if err := os.WriteFile(apiPath, []byte(staleImport), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ValidateArtifact(root, cfg); err == nil || !strings.Contains(err.Error(), "package import example does not match configuration") {
		t.Fatalf("ValidateArtifact error = %v, want stale API package import failure", err)
	}
}

func TestBuildRejectsPackageNameMismatchWithoutReplacingOutput(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	cfg := testConfig(t)
	cfg.PackageName = "differentpackage"
	cfg.OutputDir = "dist/test-package-mismatch"
	outputDir := filepath.Join(root, cfg.OutputDir)
	t.Cleanup(func() { _ = os.RemoveAll(outputDir) })
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		t.Fatal(err)
	}
	sentinel := filepath.Join(outputDir, "sentinel.txt")
	if err := os.WriteFile(sentinel, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := Build(root, cfg); err == nil || !strings.Contains(err.Error(), "does not match public Go package") {
		t.Fatalf("Build error = %v, want package name mismatch", err)
	}
	if _, err := os.Stat(sentinel); err != nil {
		t.Fatalf("existing output was replaced before configuration validation completed: %v", err)
	}
}

func TestFindUnresolvedPlaceholders(t *testing.T) {
	t.Parallel()
	directory := t.TempDir()
	if err := os.WriteFile(filepath.Join(directory, "index.html"), []byte("MODULE_PATH"), 0o644); err != nil {
		t.Fatal(err)
	}
	if failures := FindUnresolvedPlaceholders(directory); len(failures) != 1 {
		t.Fatalf("placeholder failures = %v", failures)
	}
}
