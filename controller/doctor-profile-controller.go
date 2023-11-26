package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetDoctorProfileController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	responseData, err := repository.GetDoctorProfile(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor profile",
			"response": err.Error(),
		})
	}

	doctorResponse := dto.ConvertToDoctorProfileResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor profile",
		"response": doctorResponse,
	})
}

// Work History

func GetDoctorWorkHistoriesController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	responseData, err := repository.GetDoctorWorkHistory(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor work histories",
			"response": err.Error(),
		})

	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message":  "no doctor work histories found",
			"response": nil,
		})
	}

	var doctorResponse []dto.DoctorWorkHistoryResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorWorkHistoriesResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor work histories",
		"response": doctorResponse,
	})
}

func CreateDoctorWorkHistoryController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	workHistory := dto.DoctorWorkHistoryRequest{}
	errBind := c.Bind(&workHistory)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	workData := dto.ConvertToDoctorWorkHistoryModel(workHistory)
	workData.DoctorProfileID = user

	responseData, err := repository.InsertDoctorWorkHistory(workData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create new work history",
			"response": err.Error(),
		})
	}

	workResponse := dto.ConvertToDoctorWorkHistoriesResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create new work history",
		"response": workResponse,
	})
}

func UpdateDoctorWorkHistoryController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor work history",
			"response": err.Error(),
		})
	}

	checkWork, err := repository.GetDoctorWorkHistoryByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor work history",
			"response": err.Error(),
		})
	}

	if checkWork.DoctorProfileID != user {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	workHistory := dto.DoctorWorkHistoryRequest{}
	errBind := c.Bind(&workHistory)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	workData := dto.ConvertToDoctorWorkHistoryModel(workHistory)
	workData.ID = uuid

	responseData, err := repository.UpdateDoctorWorkHistoryByID(uuid, workData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor work history",
			"response": err.Error(),
		})
	}

	responseData, err = repository.GetDoctorWorkHistoryByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor work history",
			"response": err.Error(),
		})
	}

	workResponse := dto.ConvertToDoctorWorkHistoriesResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor work history",
		"response": workResponse,
	})
}

func DeleteDoctorWorkHistoryController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor work history",
			"response": err.Error(),
		})
	}

	checkWork, err := repository.GetDoctorWorkHistoryByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor work history",
			"response": err.Error(),
		})
	}

	checkDoctorProfileID, err := repository.GetDoctorByID(checkWork.DoctorProfileID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor work history",
			"response": err.Error(),
		})
	}

	if checkDoctorProfileID.ID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor profile data.",
		})
	}

	err = repository.DeleteDoctorWorkHistoryByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor work history",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor work history",
		"response": nil,
	})
}

// Education

func GetDoctorEducationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	responseData, err := repository.GetDoctorEducation(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor educations",
			"response": err.Error(),
		})

	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message":  "no doctor education found",
			"response": nil,
		})
	}

	var doctorResponse []dto.DoctorEducationResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorEducationResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor educations",
		"response": doctorResponse,
	})
}

func CreateDoctorEducationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	education := dto.DoctorEducationRequest{}
	errBind := c.Bind(&education)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	educationData := dto.ConvertToDoctorEducationModel(education)
	educationData.DoctorProfileID = user

	responseData, err := repository.InsertDoctorEducation(educationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create doctor education",
			"response": err.Error(),
		})
	}

	educationResponse := dto.ConvertToDoctorEducationResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create doctor education",
		"response": educationResponse,
	})
}

func UpdateDoctorEducationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor education",
			"response": err.Error(),
		})
	}

	checkEducation, err := repository.GetDoctorEducationByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor education",
			"response": err.Error(),
		})
	}

	if checkEducation.DoctorProfileID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor profile data.",
		})
	}

	education := dto.DoctorEducationRequest{}
	errBind := c.Bind(&education)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	educationData := dto.ConvertToDoctorEducationModel(education)
	educationData.ID = uuid

	responseData, err := repository.UpdateDoctorEducationByID(uuid, educationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor education",
			"response": err.Error(),
		})
	}

	educationResponse := dto.ConvertToDoctorEducationResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor education",
		"response": educationResponse,
	})
}

func DeleteDoctorEducationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor education",
			"response": err.Error(),
		})
	}

	checkEducation, err := repository.GetDoctorEducationByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor education",
			"response": err.Error(),
		})
	}

	checkDoctorProfileID, err := repository.GetDoctorByID(checkEducation.DoctorProfileID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor profile",
			"response": err.Error(),
		})
	}

	if checkDoctorProfileID.ID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor profile data.",
		})
	}

	err = repository.DeleteDoctorEducationByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor education",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor education",
		"response": nil,
	})
}

// Certification

func GetDoctorCertificationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	responseData, err := repository.GetDoctorCertification(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor certification",
			"response": err.Error(),
		})

	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message":  "no doctor certifications found",
			"response": nil,
		})
	}

	var doctorResponse []dto.DoctorCertificationResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorCertificationResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor certifications",
		"response": doctorResponse,
	})
}

func CreateDoctorCertificationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	certification := dto.DoctorCertificationRequest{}
	errBind := c.Bind(&certification)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	certificationData := dto.ConvertToDoctorCertificationModel(certification)
	certificationData.DoctorProfileID = user

	responseData, err := repository.InsertDoctorCertification(certificationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create doctor certification",
			"response": err.Error(),
		})
	}

	certificationResponse := dto.ConvertToDoctorCertificationResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create doctor certification",
		"response": certificationResponse,
	})
}

func UpdateDoctorCertificationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor certification",
			"response": err.Error(),
		})
	}

	checkCertification, err := repository.GetDoctorCertificationByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor certification",
			"response": err.Error(),
		})
	}

	if checkCertification.DoctorProfileID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor profile data.",
		})
	}

	certification := dto.DoctorCertificationRequest{}
	errBind := c.Bind(&certification)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	certificationData := dto.ConvertToDoctorCertificationModel(certification)
	certificationData.ID = uuid

	responseData, err := repository.UpdateDoctorCertificationByID(uuid, certificationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor certification",
			"response": err.Error(),
		})
	}

	certificationResponse := dto.ConvertToDoctorCertificationResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor certification",
		"response": certificationResponse,
	})
}

func DeleteDoctorCertificationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor certification",
			"response": err.Error(),
		})
	}

	checkCertification, err := repository.GetDoctorCertificationByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor certification",
			"response": err.Error(),
		})
	}

	if checkCertification.DoctorProfileID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor profile data.",
		})
	}

	err = repository.DeleteDoctorCertificationByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor certification",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor certification",
		"response": "success",
	})
}