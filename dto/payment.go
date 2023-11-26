package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type PaymentRequest struct {
	Method         string     `json:"method" form:"method"`
	Name    	   string     `json:"name" form:"name"`
	AccountNumber  string     `json:"account_number" form:"account_number"`
	Image  		   string     `json:"image" form:"image"`
}

type PaymentResponse struct {
	ID    	  	   uuid.UUID  `json:"id"`
	TransactionID  uuid.UUID  `json:"transaction_id"`
	Method         string     `json:"method"`
	Name    	   string     `json:"name"`
	AccountNumber  string     `json:"account_number"`
	Image  		   string     `json:"image"`
}

func ConvertToPaymentModel(payment PaymentRequest) model.Payment {
	return model.Payment{
		ID:     	   uuid.New(),
		Method:    	   payment.Method,
		Name:  		   payment.Name,
		AccountNumber: payment.AccountNumber,
		Image: 		   payment.Image,
	}
}

func ConvertToPaymentResponse(payment model.Payment) PaymentResponse {
	return PaymentResponse{
		ID:    	   	   payment.ID,
		TransactionID: payment.TransactionID,
		Method:    	   payment.Method,
		Name:  		   payment.Name,
		AccountNumber: payment.AccountNumber,
		Image: 		   payment.Image,
	}
}
