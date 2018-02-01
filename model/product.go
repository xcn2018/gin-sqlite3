package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Name  string
	Price uint
}

func Read(keyword string) string {
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
