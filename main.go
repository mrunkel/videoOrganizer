package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Option struct {
	gorm.Model
	key string `gorm:"type:varchar(100);unique_index"`
	value string `gorm:"size:255"`
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Option{})

	// Create
	db.Create(&Option{key: "setting1", value: "one"})

	// Read
	var option Option
	db.First(&option, "setting1")                  // find option with id 1
	db.First(&option, "value = ?", "one")			 // find option with code l1212

	// Update - update option's price to 2000
	db.Model(&option).Update("value", "two")

	// Delete - delete option
	db.Delete(&option)
}
