package infrastructure

import (
	dto "Tomoya185-miyawaki/gin-todo/infrastructure/dto"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() *gorm.DB {
	config := "gin_todo:root@tcp(db:3306)/gin_todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("db init error: ", err)
	}
	autoMigrate()

	return db
}

func autoMigrate() {
	db.AutoMigrate(&dto.Todo{})
}
