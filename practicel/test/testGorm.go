package main

import (
	"fmt"
	"go-learning/practicel/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:zhb19970504@tcp(127.0.0.1:3306)/golearning"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//
	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "中转"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))
	db.Model(&user).Update("PassWord", 1234)
}
