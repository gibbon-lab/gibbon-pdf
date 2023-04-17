package main

import (
	"gibbon-lab/gibbon-pdf/internal/api"
	pdfGenerator "gibbon-lab/gibbon-pdf/internal/pdf_generator"

	"github.com/flosch/pongo2"
)

var tplSet *pongo2.TemplateSet

func main() {
	pdfGenerator.InitChrome(nil)
	defer pdfGenerator.CloseChrome()

	api.InitApi()
}
