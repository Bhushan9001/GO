package config

import (
	"fmt"

	"github.com/Bhushan9001/GO_CRUD/internal/models"
	"github.com/Bhushan9001/GO_CRUD/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(localhost:3306)/book_management?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{});

	utils.LogError(err,"Error While connecting to Database");

	err = DB.AutoMigrate(&models.User{} ,&models.Books{});

	utils.LogError(err,"Error Migrating Database");


	fmt.Printf("Database Connected!!!")

}
