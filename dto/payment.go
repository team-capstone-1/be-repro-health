package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type PaymentRequest struct {
	Name    	   string     `json:"name" form:"name"`
	AccountNumber  string     `json:"account_number" form:"account_number"`
	Image  		   string     `json:"image" form:"image"`
}

type PaymentResponse struct {
	ID    	  	   uuid.UUID  `json:"id"`
	TransactionID  uuid.UUID  `json:"transaction_id"`
	Name    	   string     `json:"name"`
	AccountNumber  string     `json:"account_number"`
	Image  		   string     `json:"image"`
}

func ConvertToPaymentModel(payment PaymentRequest) model.Payment {
	return model.Payment{
		ID:     	   uuid.New(),
		Name:  		   payment.Name,
		AccountNumber: payment.AccountNumber,
		Image: 		   payment.Image,
	}
}

func ConvertToPaymentResponse(payment model.Payment) PaymentResponse {
	return PaymentResponse{
		ID:    	   	   payment.ID,
		TransactionID: payment.TransactionID,
		Name:  		   payment.Name,
		AccountNumber: payment.AccountNumber,
		Image: 		   payment.Image,
	}
}
