package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	//gorm.Model: ID, udpated_at, craeted_at, deleted_at, etc...
	gorm.Model
	Code  string
	Name  string
	Price uint
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	initdb()

	router.GET("/", rootPath)
	router.GET("/search", searchPath)

	router.Run()
}

func initdb() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Not Found Database.")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L0001", Name: "Apple", Price: 1000})
	db.Create(&Product{Code: "L0002", Name: "Orange", Price: 100})
	db.Create(&Product{Code: "L0003", Name: "Banana", Price: 2000})
	db.Create(&Product{Code: "L0004", Name: "Papaya", Price: 5000})
	db.Create(&Product{Code: "L0005", Name: "Mango", Price: 3000})
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
