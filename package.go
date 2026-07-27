package gopackage

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	// ErrEmptyValue is returned when a label contains no non-whitespace characters.
	ErrEmptyValue = errors.New("label value is empty")
	// ErrInvalidStyle is returned when an unsupported Style is configured.
	ErrInvalidStyle = errors.New("invalid label style")
)

// Style controls how a Formatter changes the letter case of a label.
type Style string

const (
	// StylePlain preserves the case supplied by the caller.
	StylePlain Style = "plain"
	// StyleTitle uppercases the first rune of each whitespace-delimited word.
	StyleTitle Style = "title"
	// StyleUpper converts the complete label to upper case.
	StyleUpper Style = "upper"
)

// Valid reports whether the style is supported.
func (s Style) Valid() bool {
	return s == StylePlain || s == StyleTitle || s == StyleUpper
}

// Labeler is implemented by values that can turn input into a display label.
type Labeler interface {
	Label(value string) (string, error)
}

// Formatter validates and formats human-readable labels.
//
// A Formatter is immutable after construction and is safe for concurrent use.
type Formatter struct {
	prefix string
	suffix string
	style  Style
}

// Option configures a Formatter during construction.
type Option func(*Formatter) error

// New constructs a Formatter. The default formatter trims surrounding whitespace
// and otherwise preserves the caller's letter case.
func New(options ...Option) (*Formatter, error) {
	formatter := &Formatter{style: StylePlain}
	for index, option := range options {
		if option == nil {
			return nil, fmt.Errorf("option %d: nil option", index+1)
		}
		if err := option(formatter); err != nil {
			return nil, fmt.Errorf("option %d: %w", index+1, err)
		}
	}
	return formatter, nil
}

// WithAffixes returns an Option that adds text before and after each formatted label.
func WithAffixes(prefix, suffix string) Option {
	return func(formatter *Formatter) error {
		formatter.prefix = prefix
		formatter.suffix = suffix
		return nil
	}
}

// WithStyle returns an Option that selects a supported letter-case style.
func WithStyle(style Style) Option {
	return func(formatter *Formatter) error {
		if !style.Valid() {
			return fmt.Errorf("%w: %q", ErrInvalidStyle, style)
		}
		formatter.style = style
		return nil
	}
}

// Label trims, validates, and formats value according to the Formatter configuration.
func (f *Formatter) Label(value string) (string, error) {
	if f == nil {
		return "", errors.New("nil Formatter")
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return "", ErrEmptyValue
	}

	switch f.style {
	case StylePlain:
	case StyleTitle:
		value = titleWords(value)
	case StyleUpper:
		value = strings.ToUpper(value)
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidStyle, f.style)
	}

	return f.prefix + value + f.suffix, nil
}

// FormatLabel is a convenience wrapper around New and Formatter.Label.
func FormatLabel(value string, options ...Option) (string, error) {
	formatter, err := New(options...)
	if err != nil {
		return "", err
	}
	return formatter.Label(value)
}

func titleWords(value string) string {
	words := strings.Fields(value)
	for index, word := range words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		words[index] = string(runes)
	}
	return strings.Join(words, " ")
}
