package database

import (
	"capstone-project/model"
	"capstone-project/config"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
}

type DbSetup struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	database := DbSetup{
		DB_Username: config.DB_USERNAME,
		DB_Password: config.DB_PASSWORD,
		DB_Port:     config.DB_PORT,
		DB_Host:     config.DB_HOST,
		DB_Name:     config.DB_NAME,
	}
	fmt.Println("database:", database)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		database.DB_Username,
		database.DB_Password,
		database.DB_Host,
		database.DB_Port,
		database.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Patient{})
}
