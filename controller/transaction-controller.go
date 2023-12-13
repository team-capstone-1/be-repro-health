package controller

import (
	"errors"
	"net/http"

	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"capstone-project/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetTransactionController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	transactionResponse := dto.ConvertToTransactionResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transaction",
		"response": transactionResponse,
	})
}

func GetTransactionsController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	responseData, err := repository.GetTransactions(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transactions",
			"reponse": err.Error(),
		})
	}

	var transactionResponse []dto.TransactionResponse
	for _, transaction := range responseData {
		transactionResponse = append(transactionResponse, dto.ConvertToTransactionResponse(transaction))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transactions",
		"response": transactionResponse,
	})
}

func GetPatientTransactionsController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetPatientTransactions(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	var transactionResponse []dto.TransactionResponse
	for _, transaction := range responseData {
		transactionResponse = append(transactionResponse, dto.ConvertToTransactionResponse(transaction))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transaction",
		"response": transactionResponse,
	})
}

func CreatePaymentController(c echo.Context) error {
	transactionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	paymentExist := repository.CheckPayment(transactionID)
	if paymentExist {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create payment",
			"response": errors.New("Payment Already Exist").Error(),
		})
	}

	payment := dto.PaymentRequest{}
	errBind := c.Bind(&payment)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	transaction, err := repository.GetTransactionByID(transactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "transaction does not exist",
			"reponse": err.Error(),
		})
	}

	if transaction.PaymentStatus == "done" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "transaction payment status is done",
			"reponse": errors.New("Can't create Payment when Transaction Payment Status is done").Error(),
		})
	}

	paymentImage, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error upload payment image",
			"response": err.Error(),
		})
	}

	paymentImageURL, err := util.UploadToCloudinary(paymentImage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error upload profile image to Cloudinary",
			"response": err.Error(),
		})
	}

	paymentData := dto.ConvertToPaymentModel(payment)
	paymentData.Image = paymentImageURL

	paymentData.TransactionID = transactionID

	responseData, err := repository.InsertPayment(paymentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create payment",
			"response": err.Error(),
		})
	}

	paymentResponse := dto.ConvertToPaymentResponse(responseData)

	CreateNotification(
		transaction.Consultation.PatientID,
		"Pembayaran Sedang Diproses",
		"Pembayaran untuk janji temu Anda sedang dalam proses. Kami akan memberi tahu Anda begitu proses selesai.",
		"janji_temu",
	)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create new payment",
		"response": paymentResponse,
	})
}

func RescheduleController(c echo.Context) error {
	transactionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	_, err = repository.GetTransactionByID(transactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	consultation, err := repository.GetConsultationByTransactionID(transactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get consultation",
			"reponse": err.Error(),
		})
	}

	if consultation.Rescheduled {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed rescheduled",
			"reponse": errors.New("You can only reschedule once").Error(),
		})
	}

	updateData := dto.ConsultationRescheduleRequest{}

	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	doctorID := consultation.DoctorID
	dateString := updateData.Date.Format("2006-01-02")
	// Melakukan pengecekan apakah dokter sedang libur
	isDoctorOnHoliday, err := repository.IsDoctorOnHoliday(doctorID, dateString, updateData.Session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed to check doctor's holiday",
			"response": err.Error(),
		})
	}

	if isDoctorOnHoliday {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed to reschedule",
			"response": "Doctor is on holiday on the selected date and session.",
		})
	}

	rescheduleData := dto.ConvertToConsultationRescheduleModel(updateData, consultation.ID)
	rescheduleData.QueueNumber, err = repository.GenerateQueueNumber(rescheduleData.Date, rescheduleData.Session)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed generate queue number",
			"response": err.Error(),
		})
	}

	_, err = repository.RescheduleConsultation(consultation.ID, rescheduleData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed update consultation",
			"response": err.Error(),
		})
	}

	// Setelah berhasil mereschedule, perbarui status ketersediaan dokter menjadi true
	err = repository.UpdateDoctorAvailability(doctorID, true, dateString, updateData.Session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":  "failed update doctor availability",
			"response": err.Error(),
		})
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	returnData, err := repository.GetTransactionByID(transactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	transactionResponse := dto.ConvertToTransactionResponse(returnData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update consultation",
		"response": transactionResponse,
	})
}

func CancelTransactionController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	transaction, err := repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create refund",
			"reponse": err.Error(),
		})
	}

	if transaction.Consultation.PaymentMethod == "clinic_payment" {
		err := repository.UpdateTransactionStatus(uuid, "cancelled")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed cancel transaction",
				"response": err.Error(),
			})
		}

		transaction, err := repository.GetTransactionByID(uuid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "failed cancel transaction",
				"reponse": err.Error(),
			})
		}

		responseData := dto.ConvertToTransactionResponse(transaction)

		return c.JSON(http.StatusOK, map[string]any{
			"message":  "success cancel appointment",
			"response": responseData,
		})
	} else {
		paymentExist := repository.CheckPayment(uuid)
		if !paymentExist {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed create refund",
				"response": errors.New("This transaction doesn't have any payment yet").Error(),
			})
		}

		refundExist := repository.CheckRefund(uuid)
		if refundExist {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed create refund",
				"response": errors.New("Refund Already Exist").Error(),
			})
		}

		refund := dto.RefundRequest{}
		errBind := c.Bind(&refund)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "error bind data",
				"response": errBind.Error(),
			})
		}

		refundData := dto.ConvertToRefundModel(refund)
		refundData.TransactionID = uuid

		responseData, err := repository.InsertRefund(refundData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed create refund",
				"response": err.Error(),
			})
		}

		err = repository.UpdateTransactionPaymentStatus(uuid, "refund")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed create refund",
				"response": err.Error(),
			})
		}

		refundResponse := dto.ConvertToRefundResponse(responseData)

		return c.JSON(http.StatusCreated, map[string]any{
			"message":  "success create new refund",
			"response": refundResponse,
		})
	}
}

func ValidateRefund(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.UpdateRefundStatus(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update refund",
			"reponse": err.Error(),
		})
	}

	refundResponse := dto.ConvertToRefundResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update refund",
		"response": refundResponse,
	})
}

func PaymentTimeOut(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	transaction, err := repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	if transaction.PaymentStatus != "pending" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update transaction",
			"reponse": errors.New("You can't time-out a transaction if the payment status is not pending").Error(),
		})
	}

	err = repository.UpdateTransactionStatus(uuid, "cancelled")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update transaction",
			"reponse": err.Error(),
		})
	}

	transaction, err = repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	transactionResponse := dto.ConvertToTransactionResponse(transaction)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update transaction",
		"response": transactionResponse,
	})
}
