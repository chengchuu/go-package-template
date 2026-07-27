package site

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestExtractAPIUsesExportedRootPackage(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	docs, err := ExtractAPI(root, DefaultConfig().ModulePath)
	if err != nil {
		t.Fatal(err)
	}
	if docs.PackageName != "gopackage" {
		t.Fatalf("package name = %q", docs.PackageName)
	}
	want := map[string]bool{"Formatter": false, "Label": false, "FormatLabel": false, "Labeler": false, "ErrEmptyValue": false, "StylePlain": false}
	for _, item := range docs.Items {
		if _, ok := want[item.Name]; ok {
			want[item.Name] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Errorf("generated API missing %s", name)
		}
	}
}
