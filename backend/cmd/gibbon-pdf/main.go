package main

import (
	"fmt"
	"gibbon-lab/gibbon-pdf/internal/api"
	pdfGenerator "gibbon-lab/gibbon-pdf/internal/pdf_generator"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
)

var tplSet *pongo2.TemplateSet

func createTemplateSet() error {
	// Création du registre des templates
	templatePath := "../pdf-templates"
	tplSet = pongo2.NewSet("templates", pongo2.MustNewLocalFileSystemLoader(templatePath))

	// Chemin du dossier contenant les templates

	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		tplName := path[len(templatePath)+1:]
		if filepath.Ext(tplName) != ".html" {
			return nil
		}
		tplSet.FromCache(tplName)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	pdfGenerator.InitChrome(nil)
	defer pdfGenerator.CloseChrome()
	createTemplateSet()

	tpl, err := tplSet.FromCache("invoice/default.html")
	if err != nil {
		panic(err)
	}
	output, err := tpl.Execute(pongo2.Context{"total": 1000})
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

	api.InitApi()
}
