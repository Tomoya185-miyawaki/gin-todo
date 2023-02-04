package infrastructure

import (
	"fmt"
	"os"

	dto "github.com/Tomoya185-miyawaki/gin-todo/infrastructure/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func Init() *gorm.DB {
	env := os.Getenv("ENV")
	if env != "production" {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	db, err := gorm.Open("mysql", os.Getenv("DB_CONNECT"))
	if err != nil {
		fmt.Println("db init error: ", err)
	}
	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&dto.Todo{})
}
