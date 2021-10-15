package cinemasEntity

import (
	"TiBO_API/businesses/addressesEntity"
	"context"
	"time"
)

type Domain struct {
	ID          uint
	Slug        string
	Name        string
	PhoneNumber string
	IsOpen      bool
	AddressID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	CreateCinema(ctx context.Context, cinemaData *Domain, addressData *addressesEntity.Domain) (*Domain, error)
	FindByIP(ctx context.Context, ip string) ([]Domain, error)
	FindByName(ctx context.Context, name string) ([]Domain, error)
	FindBySlug(ctx context.Context, slug string) (Domain, error)
	UpdateCinema(ctx context.Context, cinemaData *Domain, addressData *addressesEntity.Domain, slug string) (*Domain, error)
	DeleteBySlug(ctx context.Context, slug string) (string, error)
}

type Repository interface {
	// Database mysql
	PostNewCinema(ctx context.Context, cinemaDomain *Domain) (Domain, error)
	GetByAddress(ctx context.Context, addressID []uint) ([]Domain, error)
	GetByName(ctx context.Context, name string) ([]Domain, error)
	GetBySlug(ctx context.Context, slug string) (Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, slugID string, cinemaDomain *Domain) (Domain, error)
	Delete(ctx context.Context, slug string) (string, error)
	FindStatusByTitle(ctx context.Context, title string) (Domain, error)
}
