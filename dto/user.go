package dto

import (
	"capstone-project/model"
)

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID 		 uint `json:"id"`
	Email    string `json:"email"`
}

func ConvertToUserModel(user UserRequest) model.User {
	return model.User{
		Email:       user.Email,
		Password:    user.Password,
	}
}

func ConvertToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Email:       user.Email,
	}
}