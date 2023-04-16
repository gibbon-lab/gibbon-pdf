package api

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func InitApi() {
	engine = gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		tpl, err := pongo2.FromString("Hello {{ name|capfirst }}!")
		if err != nil {
			panic(err)
		}
		// Now you can render the template with the given
		// pongo2.Context how often you want to.
		out, err := tpl.Execute(pongo2.Context{"name": "alex"})
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": out,
		})
	})

	engine.Run()
}
