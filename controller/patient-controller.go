package controller

import (
	"net/http"
	"sort"

	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"capstone-project/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetPatientsController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	responseData, err := repository.GetAllPatients(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get patients",
			"response": err.Error(),
		})
	}

	var patientResponse []dto.PatientResponse
	for _, patient := range responseData {
		patientResponse = append(patientResponse, dto.ConvertToPatientResponse(patient))
	}

	sort.Slice(patientResponse, func(i, j int) bool {
		return patientResponse[i].CreatedAt.Before(patientResponse[j].CreatedAt)
	})

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get patients",
		"response": patientResponse,
	})
}

func GetPatientController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetPatientByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse": err.Error(),
		})
	}

	if responseData.UserID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	patientResponse := dto.ConvertToPatientResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get patient",
		"response": patientResponse,
	})
}

func CreatePatientController(c echo.Context) error {
	patient := dto.PatientRequest{}
	errBind := c.Bind(&patient)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}
	
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}
	patientData := dto.ConvertToPatientModel(patient)

	profileImage, err := c.FormFile("profile_image")
	if err != http.ErrMissingFile{
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "error upload profile image",
				"response": err.Error(),
			})
		}
	
		profileURL, err := util.UploadToCloudinary(profileImage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "error upload profile image to Cloudinary",
				"response": err.Error(),
			})
		}
		patientData.ProfileImage = profileURL
	}
	
	patientData.UserID = user

	responseData, err := repository.InsertPatient(patientData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create patient",
			"response": err.Error(),
		})
	}

	patientResponse := dto.ConvertToPatientResponse(responseData)
	CreateNotification(
		patientResponse.ID,
		"Profil Baru Dibuat!",
		"Profil baru Anda telah berhasil dibuat. Sesuaikan profil Anda dan nikmati pengalaman kesehatan reproduksi yang lebih personal",
		"info",
	)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create new patient",
		"response": responseData,
	})
}

func UpdatePatientController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse": err.Error(),
		})
	}
	if checkPatient.UserID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	updateData := dto.PatientRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	patientData := dto.ConvertToPatientModel(updateData)

	profileImage, err := c.FormFile("profile_image")
	if err != http.ErrMissingFile{
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "error upload profile image",
				"response": err.Error(),
			})
		}
	
		profileURL, err := util.UploadToCloudinary(profileImage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "error upload profile image to Cloudinary",
				"response": err.Error(),
			})
		}
		patientData.ProfileImage = profileURL
	}
	
	patientData.ID = uuid

	responseData, err := repository.UpdatePatientByID(uuid, patientData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed update patient",
			"response": err.Error(),
		})
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	responseData, err = repository.GetPatientByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse": err.Error(),
		})
	}

	patientResponse := dto.ConvertToPatientResponse(responseData)
	CreateNotification(
		patientResponse.ID,
		"Perubahan Profil",
		"Profil Anda telah diperbarui. Pastikan semua informasi terkini dan sesuai dengan kebutuhan Anda",
		"info",
	)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update patient",
		"response": patientResponse,
	})
}

func DeletePatientController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete patient",
			"reponse": err.Error(),
		})
	}
	if checkPatient.UserID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	err = repository.DeletePatientByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed delete patient",
			"reponse": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete patient",
		"response": "success delete patient with id " + uuid.String(),
	})
}

func CreatePatientControllerTesting() echo.HandlerFunc {
	return CreatePatientController
}