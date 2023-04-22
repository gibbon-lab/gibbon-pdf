package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

var engine *gin.Engine

const templatePath = "/templates/"

func InitApi() {
	engine = gin.Default()

	engine.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "0.0.1",
		})
	})

	engine.GET("/v1/templates/", func(c *gin.Context) {
		var dirs []string
		err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() || templatePath == path {
				return nil
			}

			dirs = append(dirs, filepath.Base(path))
			return nil
		})

		if err != nil {
			fmt.Println(fmt.Errorf("computing template list: %v", err))
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"templates": dirs,
		})
	})

	engine.GET("/v1/templates/:templateName", func(c *gin.Context) {
		jsonSchemaFilePath := fmt.Sprintf(
			"%s%s/schema.json",
			templatePath,
			c.Param("templateName"),
		)
		fileBytes, err := ioutil.ReadFile(jsonSchemaFilePath)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read file"})
			return
		}

		var data interface{}
		err = json.Unmarshal(fileBytes, &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse JSON"})
			return
		}

		schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", jsonSchemaFilePath))
		m := map[string]interface{}{"type": "string"}
		documentLoader := gojsonschema.NewGoLoader(m)
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Bad body"})
			return
		}

		if result.Valid() {
			fmt.Printf("The document is valid\n")
		} else {
			fmt.Printf("The document is not valid. see errors :\n")
			for _, err := range result.Errors() {
				// Err implements the ResultError interface
				fmt.Printf("- %s\n", err)
			}
		}

		c.JSON(http.StatusOK, data)
	})

	engine.POST("/v1/templates/:templateName/html", func(c *gin.Context) {
		// show rendered html

		c.JSON(http.StatusOK, gin.H{
			// "json_schema": ,
		})
	})

	engine.POST("/v1/templates/:templateName/pdf", func(c *gin.Context) {
		// show rendered pdf

		c.JSON(http.StatusOK, gin.H{
			// "json_schema": ,
		})
	})

	engine.Run()
}
