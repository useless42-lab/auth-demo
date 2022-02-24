package models

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type PaginationData struct {
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	PageSize  int         `json:"page_size"`
	PageIndex int         `json:"page_index"`
}

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
}
