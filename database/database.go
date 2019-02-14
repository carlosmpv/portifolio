package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql compatibility
)

// Commentary commentary table
type Commentary struct {
	*gorm.Model

	Author    string `gorm:"size:31;not null"`
	Email     string `gorm:"size:31;not null"`
	Content   string `gorm:"size:255;not null"`
	Relevance int    `gorm:"default:0"`
}

// MigrateAll migrate every table
func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(&Commentary{})
}

// GetDatabase get database connection
func GetDatabase() *gorm.DB {
	fmt.Println("Loading database...")
	db, err := gorm.Open("mysql", "portifolio:OlQNcqtYduBk4u6N@/portifolio?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	fmt.Println("Database loaded!")
	return db
}
