package invoiceEntity

import (
	"context"
	"time"
)

type Domain struct {
	ID               uint
	Date             time.Time
	ShowTime         string
	TotalTicket      int
	TotalTicketPrice int
	TotalPrices      int
	AdminFee         int
	MovieID          uint
	MoviePrice       int
	MovieTitle       string
	MovieUrl         string
	UserID           uint
	CinemaID         uint
	CinemaName       string
	CinemaAddress    string
	CreatedAt        time.Time
}

type Service interface {
	Create(ctx context.Context, userId string, invoiceData *Domain) (*Domain, error)
	GetByUserID(ctx context.Context, userId string) ([]Domain, error)
}

type Repository interface {
	Create(ctx context.Context, invoiceData *Domain) (Domain, error)
	GetByUserID(ctx context.Context, userId string) ([]Domain, error)
	GetID(ctx context.Context, ID uint) ([]Domain, error)
}
