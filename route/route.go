package route

import (
	"github.com/gin-gonic/gin"
)

func DefaultRoute(router *gin.Engine) {
	router.GET("/", rootPath)
	router.GET("/search", searchPath)
}
