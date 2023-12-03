package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"capstone-project/util"
	"errors"
	"net/http"
	"strconv"

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

	responseData, err := repository.GetDoctorWorkHistories(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor work histories",
			"response": err.Error(),
		})
	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "no doctor work history found",
			"response": nil,
		})
	}

	var workHistoryResponse []dto.DoctorWorkHistoryResponse
	for _, doctor := range responseData {
		workHistoryResponse = append(workHistoryResponse, dto.ConvertToDoctorWorkHistoriesResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor work histories",
		"response": workHistoryResponse,
	})
}

func CreateDoctorWorkHistoryController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: User is not valid.",
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

	if user != user {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admins can create work histories.",
		})
	}

	if err := validateDoctorWorkHistoryRequest(workHistory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "Invalid body",
			"response": err.Error(),
		})
	}

	workData := dto.ConvertToDoctorWorkHistoryModel(workHistory)

	responseData, err := repository.InsertDoctorWorkHistory(workData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create work history",
			"response": err.Error(),
		})
	}

	workResponse := dto.ConvertToDoctorWorkHistoriesResponse(responseData)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create new work history",
		"response": workResponse,
	})
}

func UpdateDoctorWorkHistoryController(c echo.Context) error {
	workHistoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid work history ID",
			"response": "Work History ID must be a valid UUID.",
		})
	}

	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Admin is not valid.",
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

	workHistoryData, err := repository.GetDoctorWorkHistoryByID(workHistoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor work history",
			"response": err.Error(),
		})
	}

	if workHistoryData.DoctorID != workHistory.DoctorID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	workHistoryData.StartingDate = workHistory.StartingDate
	workHistoryData.EndingDate = workHistory.EndingDate
	workHistoryData.Job = workHistory.Job
	workHistoryData.Workplace = workHistory.Workplace
	workHistoryData.Position = workHistory.Position

	updatedWorkHistory, err := repository.UpdateDoctorWorkHistoryByID(workHistoryID, workHistoryData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor work history",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor work history",
		"response": updatedWorkHistory,
	})
}

func DeleteDoctorWorkHistoryController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: User is not valid.",
		})
	}
	workHistoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "invalid work history ID",
			"response": "Work History ID must be a valid UUID.",
		})
	}

	workHistory, err := repository.GetDoctorWorkHistoryByID(workHistoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor work history",
			"response": err.Error(),
		})
	}

	if workHistory.ID != workHistory.ID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to delete other user's work history.",
		})
	}

	err = repository.DeleteDoctorWorkHistoryByID(workHistoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor work history",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor work history",
		"response": "success delete doctor work history with id " + workHistoryID.String(),
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
		return c.JSON(http.StatusNotFound, map[string]any{
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
			"response": "Permission Denied: User is not valid.",
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

	if user != user {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admins can create doctor education.",
		})
	}

	if err := validateDoctorEducationRequest(education); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "Invalid body",
			"response": err.Error(),
		})
	}

	educationData := dto.ConvertToDoctorEducationModel(education)

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
	educationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid education ID",
			"response": "Education ID must be a valid UUID.",
		})
	}

	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Admin is not valid.",
		})
	}

	educationHistory := dto.DoctorEducationRequest{}
	errBind := c.Bind(&educationHistory)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	educationHistoryData, err := repository.GetDoctorEducationByID(educationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor education",
			"response": err.Error(),
		})
	}

	if educationHistoryData.DoctorID != educationHistory.DoctorID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	educationHistoryData.StartingDate = educationHistory.StartingDate
	educationHistoryData.EndingDate = educationHistory.EndingDate
	educationHistoryData.EducationProgram = educationHistory.EducationProgram
	educationHistoryData.University = educationHistory.University

	updatedEducation, err := repository.UpdateDoctorEducationByID(educationID, educationHistoryData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor education",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor education",
		"response": updatedEducation,
	})
}

func DeleteDoctorEducationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: User is not valid.",
		})
	}

	educationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid education ID",
			"response": "Education ID must be a valid UUID.",
		})
	}

	education, err := repository.GetDoctorEducationByID(educationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor education",
			"response": err.Error(),
		})
	}

	if education.ID != education.ID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to delete other user's education.",
		})
	}

	err = repository.DeleteDoctorEducationByID(educationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor education",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor education",
		"response": "success delete doctor education with ID " + educationID.String(),
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
		return c.JSON(http.StatusNotFound, map[string]any{
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
			"response": "Permission Denied: User is not valid.",
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

	if user != user {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to create certification for other user.",
		})
	}

	// if err := validateDoctorCertificationRequest(certification); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"message":  "failed create doctor certification",
	// 		"response": err.Error(),
	// 	})
	// }

	certificationData := dto.ConvertToDoctorCertificationModel(certification)

	certificationImage, err := c.FormFile("details")
	if err != http.ErrMissingFile {
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "error upload details image",
				"response": err.Error(),
			})
		}

		certificationURL, err := util.UploadToCloudinary(certificationImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message":  "error upload details image to Cloudinary",
				"response": err.Error(),
			})
		}
		size, err := getCloudinaryFileInfo(certificationURL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message":  "error getting file size from Cloudinary",
				"response": err.Error(),
			})
		}
		certificationData.FileSize = strconv.FormatInt(size, 10)
		certificationData.Details = certificationURL
	}

	responseData, err := repository.InsertDoctorCertification(certificationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create doctor certification",
			"response": err.Error(),
		})
	}

	certificationResponse := dto.ConvertToDoctorCertificationResponse(responseData)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create doctor certification",
		"response": certificationResponse,
	})
}

func getCloudinaryFileInfo(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	contentLength := resp.Header.Get("Content-Length")
	size, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func UpdateDoctorCertificationController(c echo.Context) error {
	certificationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor certification",
			"response": err.Error(),
		})
	}

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

	certificationData, err := repository.GetDoctorCertificationByID(certificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor certification",
			"response": err.Error(),
		})
	}

	if certificationData.DoctorID != certification.DoctorID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	certificationData.StartingDate = certification.StartingDate
	certificationData.EndingDate = certification.EndingDate
	certificationData.Description = certification.Description
	certificationData.CertificateType = certification.CertificateType
	// certificationData.FileSize = certification.FileSize
	certificationData.Details = certification.Details

	updateCertificated, err := repository.UpdateDoctorCertificationByID(certificationID, certificationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update doctor certification",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor certification",
		"response": updateCertificated,
	})
}

func DeleteDoctorCertificationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: User is not valid.",
		})
	}

	certificationID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor certification",
			"response": err.Error(),
		})
	}

	certification, err := repository.GetDoctorCertificationByID(certificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check doctor certification",
			"response": err.Error(),
		})
	}

	if certification.ID != certification.ID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to delete other user's certification.",
		})
	}

	err = repository.DeleteDoctorCertificationByID(certificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete doctor certification",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor certification",
		"response": "success delete doctor certification with id " + certificationID.String(),
	})
}

func validateDoctorWorkHistoryRequest(workHistory dto.DoctorWorkHistoryRequest) error {
	if workHistory.DoctorID == uuid.Nil {
		return errors.New("Doctor ID must be a valid UUID")
	}

	if _, err := uuid.Parse(workHistory.DoctorID.String()); err != nil {
		return errors.New("Doctor ID must be a valid UUID")
	}

	if workHistory.DoctorID == uuid.Nil || workHistory.StartingDate.IsZero() || workHistory.EndingDate.IsZero() || workHistory.Job == "" || workHistory.Workplace == "" || workHistory.Position == "" {
		return errors.New("All fields must be filled in")
	}

	return nil
}

func validateDoctorEducationRequest(education dto.DoctorEducationRequest) error {
	if education.DoctorID == uuid.Nil {
		return errors.New("Doctor ID must be a valid UUID")
	}

	if _, err := uuid.Parse(education.DoctorID.String()); err != nil {
		return errors.New("Doctor ID must be a valid UUID")
	}

	if education.DoctorID == uuid.Nil || education.StartingDate.IsZero() || education.EndingDate.IsZero() || education.EducationProgram == "" || education.University == "" {
		return errors.New("All fields must be filled in")
	}

	return nil
}

// func validateDoctorCertificationRequest(certification dto.DoctorCertificationRequest) error {
// 	if certification.DoctorID == uuid.Nil {
// 		return errors.New("Doctor ID must be a valid UUID")
// 	}

// 	if _, err := uuid.Parse(certification.DoctorID.String()); err != nil {
// 		return errors.New("Doctor ID must be a valid UUID")
// 	}

// 	if certification.DoctorID == uuid.Nil || certification.StartingDate.IsZero() || certification.EndingDate.IsZero() || certification.Description == "" || certification.CertificateType == "" || certification.Details == "" {
// 		return errors.New("All fields must be filled in")
// 	}

// 	return nil
// }
