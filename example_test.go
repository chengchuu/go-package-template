package gopackage_test

import (
	"fmt"

	gopackage "github.com/chengchuu/go-package-template"
)

func ExampleFormatLabel() {
	label, err := gopackage.FormatLabel(
		"go package",
		gopackage.WithStyle(gopackage.StyleTitle),
		gopackage.WithAffixes("[", "]"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(label)
	// Output: [Go Package]
}
