package database

import (
	"capstone-project/config"
	"capstone-project/model"

	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
	AdminSeeders()
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
	DB.AutoMigrate(&model.User{}, &model.Doctor{})
	DB.AutoMigrate(&model.User{}, &model.Doctor{})
	DB.AutoMigrate(&model.Patient{})
	DB.AutoMigrate(&model.Forum{})
	DB.AutoMigrate(&model.Specialist{})
	DB.AutoMigrate(&model.Forum{})
}

func AdminSeeders() {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("Admin@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}
	admin := []model.User{
		{
			ID:       uuid.New(),
			Email:    "admin1@gmail.com",
			Password: string(passwordHash),
		},
		{
			ID:       uuid.New(),
			Email:    "admin2@gmail.com",
			Password: string(passwordHash),
		},
		{
			ID:       uuid.New(),
			Email:    "admin3@gmail.com",
			Password: string(passwordHash),
		},
	}

	for _, v := range admin {
		var exist model.User

		errCheck := DB.Where("email = ?", v.Email).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
	DB.AutoMigrate(&model.Forum{})
}

func AdminSeeders() {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("Admin@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}
	admin := []model.User{
		{
			ID:       uuid.New(),
			Email:    "admin1@gmail.com",
			Password: string(passwordHash),
		},
		{
			ID:       uuid.New(),
			Email:    "admin2@gmail.com",
			Password: string(passwordHash),
		},
		{
			ID:       uuid.New(),
			Email:    "admin3@gmail.com",
			Password: string(passwordHash),
		},
	}

	for _, v := range admin {
		var exist model.User

		errCheck := DB.Where("email = ?", v.Email).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}
