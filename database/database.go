package database

import (
	"capstone-project/config"
	"capstone-project/model"
	"time"

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
	// Seeders()
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
	DB.AutoMigrate(&model.User{}, &model.Doctor{}, &model.DoctorWorkHistory{}, &model.DoctorEducation{}, &model.DoctorCertification{})
	DB.AutoMigrate(&model.Patient{})
	DB.AutoMigrate(&model.Article{})
	DB.AutoMigrate(&model.Specialist{})
	DB.AutoMigrate(&model.Consultation{})
	DB.AutoMigrate(&model.Transaction{})
	DB.AutoMigrate(&model.Payment{})
	DB.AutoMigrate(&model.Refund{})
	DB.AutoMigrate(&model.Forum{})
	DB.AutoMigrate(&model.ForumReply{})
	DB.AutoMigrate(&model.Comment{})
}

func Seeders() {
	// ADMIN SEEDERS
	adminPasswordHash, err := bcrypt.GenerateFromPassword([]byte("Admin@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}
	admin := []model.User{
		{
			ID:       uuid.New(),
			Email:    "admin1@gmail.com",
			Password: string(adminPasswordHash),
		},
		{
			ID:       uuid.New(),
			Email:    "admin2@gmail.com",
			Password: string(adminPasswordHash),
		},
		{
			ID:       uuid.New(),
			Email:    "admin3@gmail.com",
			Password: string(adminPasswordHash),
		},
	}

	for _, v := range admin {
		var exist model.User

		errCheck := DB.Where("email = ?", v.Email).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}

	specialistID, err := uuid.Parse("0a8f12c6-1244-43e6-a704-126b173a8732")
	if err != nil {

		return
	}

	clinicID, err := uuid.Parse("fd9d3738-3bcf-4693-9079-57ad3b375af5")
	if err != nil {
		return
	}

	doctorID, err := uuid.Parse("f7613c10-29fd-4b82-bfea-1649ae41af98")
	if err != nil {
		return
	}

	// CLINIC SEEDERS
	clinic := []model.Clinic{
		{
			ID:       clinicID,
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

	// SPECIALIST SEEDERS
	specialist := []model.Specialist{
		{
			ID:    specialistID,
			Name:  "Kandungan",
			Image: "https://res.cloudinary.com/dw3n2ondc/image/upload/v1700464911/Reproduction-Health/rrau8asadj80uuugksdj.png",
		},
	}

	for _, v := range specialist {
		var exist model.Specialist

		errCheck := DB.Where("name = ?", v.Name).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}

	// DOCTOR SEEDERS
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("Doctor@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	doctor := []model.Doctor{
		{
			ID:           doctorID,
			Name:         "Dr. Andi Cahaya, Sp.OG",
			Email:        "andicahyo@gmail.com",
			Password:     string(passwordHash),
			Price:        150000,
			Address:      "Klinik Nasional. Jl. Bedrek No.47e, Sanggrahan, Condongcatur, Kec. Depok, Kabupaten Sleman, DIY",
			Phone:        "+62 812345865",
			ProfileImage: "https://res.cloudinary.com/dw3n2ondc/image/upload/v1700464491/Reproduction-Health/kypalfa66lmowdu5zh9p.png",
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

	// DOCTOR WORK HISTORIES SEEDERS
	workHistory := []model.DoctorWorkHistory{
		{
			ID:              uuid.New(),
			DoctorID: doctorID,
			StartingDate:    time.Date(2020, 07, 27, 0, 0, 0, 0, time.UTC),
			EndingDate:      time.Date(2025, 07, 27, 0, 0, 0, 0, time.UTC),
			Job:             "Konsultan Kesehatan Reproduksi",
			Workplace:       "Klinik Sehat Hati",
			Position:        "Memberikan konsultasi kepada pasien tentang kesehatan reproduksi.",
		},
		{
			ID:              uuid.New(),
			DoctorID: doctorID,
			StartingDate:    time.Date(2016, 07, 27, 0, 0, 0, 0, time.UTC),
			EndingDate:      time.Date(2019, 07, 27, 0, 0, 0, 0, time.UTC),
			Job:             "Spesialis Obstetri dan Ginekologi",
			Workplace:       "Rumah Sakit Kharisma",
			Position:        "Memperoleh gelar spesialis dalam Obstetri dan Ginekologi (Sp.OG)",
		},
	}

	for _, v := range workHistory {
		var exist model.DoctorWorkHistory

		errCheck := DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}

	// DOCTOR EDUCATIONS SEEDERS
	educations := []model.DoctorEducation{
		{
			ID:               uuid.New(),
			DoctorID:  doctorID,
			StartingDate:     time.Date(2013, 07, 27, 0, 0, 0, 0, time.UTC),
			EndingDate:       time.Date(2015, 07, 27, 0, 0, 0, 0, time.UTC),
			EducationProgram: "Program Magister Kedokteran",
			University:       "Universitas Gadjah Mada",
		},
		{
			ID:               uuid.New(),
			DoctorID:  doctorID,
			StartingDate:     time.Date(2009, 07, 27, 0, 0, 0, 0, time.UTC),
			EndingDate:       time.Date(2013, 07, 27, 0, 0, 0, 0, time.UTC),
			EducationProgram: "Program Residen Dokter Spesialis",
			University:       "Rumah Sakit Sejahtera",
		},
	}

	for _, v := range educations {
		var exist model.DoctorEducation

		errCheck := DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}

	certifications := []model.DoctorCertification{
		{
			ID:              uuid.New(),
			DoctorID: doctorID,
			CertificateType: "Sertifikasi Lisensi",
			Description:     "Praktik Medis",
			StartingDate:    time.Date(2022, 07, 27, 0, 0, 0, 0, time.UTC),
			EndingDate:      time.Date(2027, 07, 27, 0, 0, 0, 0, time.UTC),
			FileSize:        "5MB",
			Details:         "https://res.cloudinary.com/dw3n2ondc/image/upload/v1700466108/Reproduction-Health/ickckqmok4hbajzdkhpx.png",
		},
		{
			ID:              uuid.New(),
			DoctorID: doctorID,
			CertificateType: "Sertifikasi Lisensi",
			Description:     "Praktik Medis",
			StartingDate:    time.Date(2022, 07, 27, 0, 0, 0, 0, time.UTC),
			EndingDate:      time.Date(2027, 07, 27, 0, 0, 0, 0, time.UTC),
			FileSize:        "5MB",
			Details:         "https://res.cloudinary.com/dw3n2ondc/image/upload/v1700466108/Reproduction-Health/ickckqmok4hbajzdkhpx.png",
		},
	}

	for _, v := range certifications {
		var exist model.DoctorCertification

		errCheck := DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}
