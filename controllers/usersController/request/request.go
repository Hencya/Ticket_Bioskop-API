package request

import (
	"TiBO_API/businesses/usersEntity"

	"github.com/google/uuid"
)

type UserRegistration struct {
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password" validate:"required,password"`
}

type UserLogin struct {
	Uuid     uuid.UUID `json:"uuid"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,password"`
}

type UserUploadAvatar struct {
	Avatar string `json:"avatar"`
}

func (rec *UserRegistration) ToDomain() *usersEntity.Domain {
	return &usersEntity.Domain{
		Name:        rec.Name,
		Username:    rec.Username,
		Email:       rec.Email,
		PhoneNumber: rec.PhoneNumber,
		Avatar:      rec.Avatar,
		Password:    rec.Password,
	}
}
