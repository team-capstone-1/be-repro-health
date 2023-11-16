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
}

func ClinicSeeders() {
	clinic := []model.Clinic{
		{
			ID:       uuid.New(),
			Name:     "Klinik Inter Medika",
			City:     "Jakarta Pusat",
			Location: "Jl. K.S. Tubun No.21, Petamburan, Kota Jakarta Pusat ",
			Profile:  "Selamat datang di Klinik Inter Medika, tempat di mana kesehatan dan kenyamanan pasien menjadi prioritas utama. Klinik kami berkomitmen untuk memberikan pelayanan medis berkualitas tinggi.",
		},
	}

	for _, v := range clinic {
		var exist model.Clinic

		errCheck := DB.Where("name = ?", v.Name).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}

func SpecialistSeeders() {
	specialist := []model.Specialist{
		{
			ID:    uuid.New(),
			Name:  "Kandungan",
			Image: "",
		},
	}

	for _, v := range specialist {
		var exist model.Specialist

		errCheck := DB.Where("name = ?", v.Name).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}

func DoctorSeeders() {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("Doctor@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	specialistID, err := uuid.Parse("9d33c1eb-e918-4df2-bc41-b54e2fc4002c")
	if err != nil {
		// Handle error
		return
	}

	// Objek Clinic sudah ada atau dibuat
	clinicID, err := uuid.Parse("fff50db6-ca43-46fb-8d64-fa0c46f429a7")
	if err != nil {
		// Handle error
		return
	}

	doctor := []model.Doctor{
		{
			ID:           uuid.New(),
			Name:         "Dr. Andi Cahaya, Sp.OG",
			Email:        "andicahyo@gmail.com",
			Password:     string(passwordHash),
			Price:        150000,
			Address:      "Klinik Nasional. Jl. Bedrek No.47e, Sanggrahan, Condongcatur, Kec. Depok, Kabupaten Sleman, DIY",
			Phone:        "+62 812345865",
			SpecialistID: specialistID,
			ClinicID:     clinicID,
		},
	}

	for _, v := range doctor {
		var exist model.Doctor

		errCheck := DB.Where("email = ?", v.Email).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}
