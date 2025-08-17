package main

import (
	"os"
	"text/template"

	"Stella.Protostar/internal"
)

func main() {
	flags := internal.ParseFlags()

	// read template file
	templatePath := "templates/changelog_template.md"
	tmpl := template.Must(template.ParseFiles(templatePath))
	tmpl.ExecuteTemplate(os.Stdout, "changelog_template.md", flags)
}
