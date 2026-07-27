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
	cfg := DefaultConfig()
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
