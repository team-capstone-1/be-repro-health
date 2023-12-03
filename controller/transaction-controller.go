package controller

import (
	"errors"
	"net/http"

	"capstone-project/dto"
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
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	paymentExist := repository.CheckPayment(uuid)
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

	transaction, err := repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create payment",
			"reponse": err.Error(),
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
	paymentData.TransactionID = uuid

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

	updateData := dto.ConsultationRescheduleRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	rescheduleData := dto.ConvertToConsultationRescheduleModel(updateData, consultation.ID)

	_, err = repository.RescheduleConsultation(consultation.ID, rescheduleData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed update consultation",
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

	if transaction.Payment.Method == "clinic_payment" {
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

		return c.JSON(http.StatusOK, map[string]any{
			"message":  "success cancel appointment",
			"response": transaction,
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
