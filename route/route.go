package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Name  string
	Price uint
}

func DefaultRoute(router *gin.Engine) {
	router.GET("/", rootPath)
	router.GET("/search", searchPath)
}

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
	info = read(keyword)
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Result", "info": info},
	)
}

func read(keyword string) string {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Not Found Database")
	}

	var product Product
	db.First(&product, "Name = ?", keyword)

	var info string
	info += "Name: " + fmt.Sprint(product.Name)
	info += " Price: " + fmt.Sprint(product.Price)

	return info
}
