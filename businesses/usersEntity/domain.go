package usersEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Domain struct {
	ID          uint
	Uuid        uuid.UUID
	Name        string
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
	Avatar      string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	Register(ctx context.Context, data *Domain) (*Domain, error)
	Login(ctx context.Context, email string, password string) (string, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	Logout(ctx echo.Context) error
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}

type Repository interface {
	// Databases mysql
	CreateNewUser(ctx context.Context, data *Domain) (*Domain, error)
	UpdateUser(ctx context.Context, id string, data *Domain) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	DeleteUserByUuid(ctx context.Context, id string) (string, error)
}
