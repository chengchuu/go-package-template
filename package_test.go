package gopackage_test

import (
	"errors"
	"testing"

	gopackage "github.com/chengchuu/go-package-template"
)

func TestFormatLabel(t *testing.T) {
	t.Parallel()

	got, err := gopackage.FormatLabel("  go package  ",
		gopackage.WithStyle(gopackage.StyleTitle),
		gopackage.WithAffixes("[", "]"),
	)
	if err != nil {
		t.Fatalf("FormatLabel returned an error: %v", err)
	}
	if want := "[Go Package]"; got != want {
		t.Fatalf("FormatLabel() = %q, want %q", got, want)
	}
}

func TestFormatLabelRejectsEmptyInput(t *testing.T) {
	t.Parallel()

	_, err := gopackage.FormatLabel(" \t ")
	if !errors.Is(err, gopackage.ErrEmptyValue) {
		t.Fatalf("FormatLabel error = %v, want ErrEmptyValue", err)
	}
}

func TestNewRejectsInvalidStyle(t *testing.T) {
	t.Parallel()

	_, err := gopackage.New(gopackage.WithStyle("sentence"))
	if !errors.Is(err, gopackage.ErrInvalidStyle) {
		t.Fatalf("New error = %v, want ErrInvalidStyle", err)
	}
}

func TestFormatterImplementsLabeler(t *testing.T) {
	t.Parallel()

	formatter, err := gopackage.New(gopackage.WithStyle(gopackage.StyleUpper))
	if err != nil {
		t.Fatal(err)
	}
	var labeler gopackage.Labeler = formatter
	got, err := labeler.Label("release candidate")
	if err != nil {
		t.Fatal(err)
	}
	if got != "RELEASE CANDIDATE" {
		t.Fatalf("Label() = %q", got)
	}
}
