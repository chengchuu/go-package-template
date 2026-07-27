package site

import "testing"

func TestDefaultConfigDerivations(t *testing.T) {
	t.Parallel()
	cfg := DefaultConfig()
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
	if cfg.PagesBasePath != "/go-package-template/" {
		t.Fatalf("PagesBasePath = %q", cfg.PagesBasePath)
	}
	if cfg.PagesURL != "https://chengchuu.github.io/go-package-template/" {
		t.Fatalf("PagesURL = %q", cfg.PagesURL)
	}
	if cfg.Theme.StorageKey == "mazey-npm-template-theme" || cfg.CachePrefix == "mazey-npm-template-site-" {
		t.Fatal("target identity must not reuse reference PWA keys")
	}
}

func TestDefaultThemePalette(t *testing.T) {
	t.Parallel()
	cfg := DefaultConfig()
	want := ThemeConfig{
		StorageKey: "go-package-template-theme",
		ColorLight: "#f7f8fc",
		ColorDark:  "#0d1220",
		Primary: ThemePalette{
			Light: PrimaryPalette{Base: "#5b3fd6", Hover: "#4229b5", Active: "#362097", Soft: "#ece8ff", RGB: "91, 63, 214", HoverRGB: "66, 41, 181"},
			Dark:  PrimaryPalette{Base: "#a997ff", Hover: "#c3b7ff", Active: "#d9d2ff", Soft: "#29234c", RGB: "169, 151, 255", HoverRGB: "195, 183, 255"},
		},
	}
	if cfg.Theme != want {
		t.Fatalf("theme = %+v, want %+v", cfg.Theme, want)
	}
}

func TestDerivePagesBasePath(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"library":        "/library/",
		" team library ": "/team%20library/",
		"":               "/",
		"../unsafe":      "/",
	}
	for input, want := range tests {
		if got := DerivePagesBasePath(input); got != want {
			t.Errorf("DerivePagesBasePath(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestCanonicalURLs(t *testing.T) {
	t.Parallel()
	cfg := DefaultConfig()
	want := map[string]string{
		"home":     "https://chengchuu.github.io/go-package-template/",
		"examples": "https://chengchuu.github.io/go-package-template/examples/",
		"api":      "https://chengchuu.github.io/go-package-template/api/",
	}
	for name, expected := range want {
		page, ok := cfg.Page(name)
		if !ok {
			t.Fatalf("page %q not found", name)
		}
		if got := cfg.CanonicalURL(page); got != expected {
			t.Errorf("canonical %s = %q, want %q", name, got, expected)
		}
	}
}
