package response

import (
	"TiBO_API/businesses/usersEntity"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromDomain(domain *usersEntity.Domain) *Users {
	return &Users{
		Uuid:        domain.Uuid,
		Name:        domain.Name,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		Avatar:      domain.Avatar,
		CreatedAt:   domain.CreatedAt,
	}
}

type UserLogin struct {
	Token string `json:"token"`
}
