package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type OTPRequest struct {
	Email    string `json:"email" form:"email"`
}

type ValidateOTPRequest struct {
	Email    string `json:"email" form:"email"`
	OTP      string `json:"otp" form:"otp"`
}

type ChangeUserPasswordRequest struct {
	ID	     uuid.UUID `json:"id" form:"id"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func ConvertToUserModel(user UserRequest) model.User {
	return model.User{
		ID:       uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ConvertToChangeUserPasswordModel(user ChangeUserPasswordRequest) model.User {
	return model.User{
		ID:       user.ID,
		Password: user.Password,
	}
}

func ConvertToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
