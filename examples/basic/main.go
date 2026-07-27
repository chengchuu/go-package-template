// Command basic demonstrates the public label-formatting API.
package main

import (
	"fmt"
	"log"

	gopackage "github.com/chengchuu/go-package-template"
)

func main() {
	label, err := gopackage.FormatLabel(
		"go package",
		gopackage.WithStyle(gopackage.StyleTitle),
		gopackage.WithAffixes("[", "]"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(label)
}
