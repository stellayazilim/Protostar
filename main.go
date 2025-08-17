package main

import (
	"log"
	"os"
	"text/template"

	"Stella.Protostar/internal"
)

func main() {
	flags := internal.ParseFlags()

	// get template path from flags or default
	templatePath := flags["template"]
	if templatePath == "" {
		templatePath = "templates/changelog_template.md"
	}

	// parse template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	// execute template to stdout
	if err := tmpl.Execute(os.Stdout, flags); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}
