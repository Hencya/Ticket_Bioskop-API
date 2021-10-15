package usersRepo

import (
	"TiBO_API/businesses/usersEntity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uint      `gorm:"primary_key:auto_increment"`
	Uuid        uuid.UUID `gorm:"type:varchar(255)"`
	Name        string    `gorm:"type:varchar(255)"`
	Username    string    `gorm:"type:varchar(50)"`
	Email       string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password    string    `gorm:"->;<-;not null" `
	PhoneNumber string    `gorm:"type:varchar(15)"`
	Role        string    `gorm:"type:varchar(5)"`
	Avatar      string    `gorm:"type:varchar(255)"`
}

func ToDomain(rec *Users) usersEntity.Domain {
	return usersEntity.Domain{
		ID:          rec.ID,
		Uuid:        rec.Uuid,
		Name:        rec.Name,
		Username:    rec.Username,
		Email:       rec.Email,
		Password:    rec.Password,
		PhoneNumber: rec.PhoneNumber,
		Role:        rec.Role,
		Avatar:      rec.Avatar,
	}
}

func FromDomain(userDomain *usersEntity.Domain) *Users {
	return &Users{
		ID:          userDomain.ID,
		Uuid:        userDomain.Uuid,
		Name:        userDomain.Name,
		Username:    userDomain.Username,
		Email:       userDomain.Email,
		Password:    userDomain.Password,
		PhoneNumber: userDomain.PhoneNumber,
		Role:        userDomain.Role,
		Avatar:      userDomain.Avatar,
	}
}
