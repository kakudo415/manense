package main

import (
	"./page"
	"github.com/gin-gonic/gin"
)

var app = gin.New()

func main() {
	app.GET("/", page.Index)
	app.Run("127.0.0.1:8000")
}

func init() {
	app.LoadHTMLGlob("view/*.html")
}
