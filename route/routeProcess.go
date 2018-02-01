package route

import (
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

func rootPath(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Main website", "info": ""},
	)
}

func searchPath(c *gin.Context) {
	var keyword string
	keyword = c.Query("keyword")

	var info string
	info = model.Read(keyword)
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Result", "info": info},
	)
}
