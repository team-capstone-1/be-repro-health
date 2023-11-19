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
	ClinicSeeders()
	SpecialistSeeders()
	DoctorSeeders()
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
	DB.AutoMigrate(&model.Doctor{})
	DB.AutoMigrate(&model.Article{})
	DB.AutoMigrate(&model.Clinic{})
	DB.AutoMigrate(&model.Specialist{})
	DB.AutoMigrate(&model.Patient{})
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("Andi@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	specialistID, err := uuid.Parse("f8286c0b-a33b-43fa-b27e-0f87008b5349")
	if err != nil {
		// Handle error
		return
	}

	// Objek Clinic sudah ada atau dibuat
	clinicID, err := uuid.Parse("75d53030-0fe7-466a-8dce-0d481c4c1fab")
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
