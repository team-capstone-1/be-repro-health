package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func ConvertToUserModel(user UserRequest) model.User {
	return model.User{
		ID:       uuid.New(),
		Email:    user.Email,
		Password: user.Password,
	}
}

func ConvertToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}
