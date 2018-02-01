package main

import (
	"./route"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	route.DefaultRoute(router)
	router.Run()
}
