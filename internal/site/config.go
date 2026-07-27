package site

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

// SiteConfig is the build-time source of truth for project, website, SEO, theme, and PWA identity.
// It belongs to internal tooling and is not part of the public package API.
type SiteConfig struct {
	PackageName     string
	DisplayName     string
	ModulePath      string
	Description     string
	RepositoryOwner string
	RepositoryName  string
	RepositoryURL   string
	PagesBasePath   string
	PagesURL        string
	Theme           ThemeConfig
	CachePrefix     string
	PWAShortName    string
	GoVersion       string
	License         string
	OutputDir       string
	SiteSourceDir   string
	Favicon         string
	Icon192         string
	Icon512         string
	MaskableIcon512 string
	AppleTouchIcon  string
	OpenGraphImage  string
	Pages           []PageConfig
}

// ThemeConfig owns the website color modes and persisted preference identity.
type ThemeConfig struct {
	StorageKey string
	ColorLight string
	ColorDark  string
	Primary    ThemePalette
}

// ThemePalette contains the coordinated primary colors for both Bootstrap modes.
type ThemePalette struct {
	Light PrimaryPalette
	Dark  PrimaryPalette
}

// PrimaryPalette defines interactive and soft-accent colors for one mode.
type PrimaryPalette struct {
	Base     string
	Hover    string
	Active   string
	Soft     string
	RGB      string
	HoverRGB string
}

// PageConfig owns the unique identity and production location of one public page.
type PageConfig struct {
	Name        string
	OutputPath  string
	Path        string
	Title       string
	Description string
	SchemaType  string
}

// DefaultConfig returns the only editable project identity configuration.
func DefaultConfig() SiteConfig {
	const (
		owner = "chengchuu"
		repo  = "go-package-template"
	)
	basePath := DerivePagesBasePath(repo)
	pagesURL := ProductionURL(owner, repo)
	return SiteConfig{
		PackageName:     "gopackage",
		DisplayName:     "Go Package Template",
		ModulePath:      "github.com/chengchuu/go-package-template",
		Description:     "A production-ready starting point for reusable Go packages with documentation, examples, and a public website.",
		RepositoryOwner: owner,
		RepositoryName:  repo,
		RepositoryURL:   "https://github.com/chengchuu/go-package-template",
		PagesBasePath:   basePath,
		PagesURL:        pagesURL,
		Theme: ThemeConfig{
			StorageKey: "go-package-template-theme",
			ColorLight: "#f7f8fc",
			ColorDark:  "#0d1220",
			Primary: ThemePalette{
				Light: PrimaryPalette{Base: "#5b3fd6", Hover: "#4229b5", Active: "#362097", Soft: "#ece8ff", RGB: "91, 63, 214", HoverRGB: "66, 41, 181"},
				Dark:  PrimaryPalette{Base: "#a997ff", Hover: "#c3b7ff", Active: "#d9d2ff", Soft: "#29234c", RGB: "169, 151, 255", HoverRGB: "195, 183, 255"},
			},
		},
		CachePrefix:     "go-package-template-pages-",
		PWAShortName:    "Go Package",
		GoVersion:       "1.25",
		License:         "MIT",
		OutputDir:       "dist/pages",
		SiteSourceDir:   "site",
		Favicon:         "logo-purple-circle-transparent-32x32.png",
		Icon192:         "logo-purple-circle-transparent-192x192.png",
		Icon512:         "logo-purple-circle-transparent-512x512.png",
		MaskableIcon512: "logo-purple-circle-transparent-maskable-512x512.png",
		AppleTouchIcon:  "logo-purple-circle-transparent-192x192.png",
		OpenGraphImage:  "logo-purple-circle-open-graph-1200x630.png",
		Pages: []PageConfig{
			{Name: "home", OutputPath: "index.html", Path: "", Title: "Go Package Template | Production-ready Go library starter", Description: "Build a reusable Go package with tested APIs, executable examples, generated documentation, SEO, PWA support, and GitHub Pages.", SchemaType: "SoftwareSourceCode"},
			{Name: "examples", OutputPath: "examples/index.html", Path: "examples/", Title: "Go examples | Go Package Template", Description: "Read executable, source-derived examples for installing and using the Go Package Template public API.", SchemaType: "TechArticle"},
			{Name: "api", OutputPath: "api/index.html", Path: "api/", Title: "Go API reference | Go Package Template", Description: "Generated reference documentation for the exported Go Package Template package API, including functions, types, methods, constants, and errors.", SchemaType: "TechArticle"},
		},
	}
}

// DerivePagesBasePath derives a GitHub project Pages path from a repository name.
func DerivePagesBasePath(repositoryName string) string {
	name := strings.Trim(strings.TrimSpace(repositoryName), "/")
	if name == "" || strings.Contains(name, "..") {
		return "/"
	}
	return "/" + url.PathEscape(name) + "/"
}

// ProductionURL derives the canonical GitHub Pages URL for a project repository.
func ProductionURL(owner, repositoryName string) string {
	owner = strings.TrimSpace(owner)
	base := DerivePagesBasePath(repositoryName)
	return fmt.Sprintf("https://%s.github.io%s", owner, base)
}

// CanonicalURL returns the absolute production URL for page.
func (c SiteConfig) CanonicalURL(page PageConfig) string {
	return c.PagesURL + page.Path
}

// AssetURL returns a repository-subpath-aware URL for a generated asset.
func (c SiteConfig) AssetURL(assetPath string) string {
	cleaned := strings.TrimPrefix(path.Clean("/"+assetPath), "/")
	if strings.HasSuffix(assetPath, "/") && cleaned != "" {
		cleaned += "/"
	}
	return c.PagesBasePath + cleaned
}

// Page finds a configured page by name.
func (c SiteConfig) Page(name string) (PageConfig, bool) {
	for _, page := range c.Pages {
		if page.Name == name {
			return page, true
		}
	}
	return PageConfig{}, false
}

// Validate checks invariants needed by generation and deployment.
func (c SiteConfig) Validate() error {
	required := map[string]string{
		"package name": c.PackageName, "display name": c.DisplayName, "module path": c.ModulePath,
		"description": c.Description, "repository owner": c.RepositoryOwner, "repository name": c.RepositoryName,
		"repository URL": c.RepositoryURL, "Pages URL": c.PagesURL, "theme storage key": c.Theme.StorageKey,
		"cache prefix": c.CachePrefix, "PWA short name": c.PWAShortName, "favicon": c.Favicon,
		"192px icon": c.Icon192, "512px icon": c.Icon512, "maskable icon": c.MaskableIcon512,
		"Apple touch icon": c.AppleTouchIcon, "Open Graph image": c.OpenGraphImage,
		"light background": c.Theme.ColorLight, "dark background": c.Theme.ColorDark,
		"light primary": c.Theme.Primary.Light.Base, "light primary hover": c.Theme.Primary.Light.Hover,
		"light primary active": c.Theme.Primary.Light.Active, "light primary soft": c.Theme.Primary.Light.Soft,
		"light primary RGB": c.Theme.Primary.Light.RGB, "light primary hover RGB": c.Theme.Primary.Light.HoverRGB,
		"dark primary": c.Theme.Primary.Dark.Base, "dark primary hover": c.Theme.Primary.Dark.Hover,
		"dark primary active": c.Theme.Primary.Dark.Active, "dark primary soft": c.Theme.Primary.Dark.Soft,
		"dark primary RGB": c.Theme.Primary.Dark.RGB, "dark primary hover RGB": c.Theme.Primary.Dark.HoverRGB,
	}
	for label, value := range required {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s must not be empty", label)
		}
	}
	if c.PagesBasePath != DerivePagesBasePath(c.RepositoryName) {
		return fmt.Errorf("Pages base path %q does not match repository %q", c.PagesBasePath, c.RepositoryName)
	}
	if c.PagesURL != ProductionURL(c.RepositoryOwner, c.RepositoryName) {
		return fmt.Errorf("Pages URL %q is not derived from repository identity", c.PagesURL)
	}
	if len(c.Pages) != 3 {
		return fmt.Errorf("exactly three stable public pages are required")
	}
	seen := make(map[string]bool)
	for _, page := range c.Pages {
		if page.Name == "" || page.OutputPath == "" || page.Title == "" || page.Description == "" {
			return fmt.Errorf("page configuration is incomplete: %+v", page)
		}
		if seen[page.OutputPath] {
			return fmt.Errorf("duplicate page output path %q", page.OutputPath)
		}
		seen[page.OutputPath] = true
	}
	return nil
}
