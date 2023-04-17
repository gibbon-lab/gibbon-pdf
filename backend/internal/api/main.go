package api

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

const templatePath = "../pdf-templates/"

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

		c.JSON(200, gin.H{
			"templates": dirs,
		})
	})

	engine.Run()
}
