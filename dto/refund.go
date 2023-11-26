package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type RefundRequest struct {
	Name    	   string     `json:"name" form:"name"`
	AccountNumber  string     `json:"account_number" form:"account_number"`
}

type RefundResponse struct {
	ID    	  	   uuid.UUID  `json:"id"`
	TransactionID  uuid.UUID  `json:"transaction_id"`
	Name    	   string     `json:"name"`
	AccountNumber  string     `json:"account_number"`
	Date  		   time.Time  `json:"date"`
	Status    	   string     `json:"status"`
}

func ConvertToRefundModel(refund RefundRequest) model.Refund {
	return model.Refund{
		ID:     	   uuid.New(),
		Name:  		   refund.Name,
		AccountNumber: refund.AccountNumber,
		Date: 		   time.Now(),
		Status: 	   model.RefundStatus(model.Processing),
	}
}

func ConvertToRefundResponse(refund model.Refund) RefundResponse {
	return RefundResponse{
		ID:    	   	   refund.ID,
		TransactionID: refund.TransactionID,
		Name:  		   refund.Name,
		AccountNumber: refund.AccountNumber,
		Date: 		   refund.Date,
		Status: 	   string(refund.Status),
	}
}
