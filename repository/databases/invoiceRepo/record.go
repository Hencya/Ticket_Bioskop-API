package invoiceRepo

import (
	"TiBO_API/businesses/invoiceEntity"
	"TiBO_API/repository/databases/cinemasRepo"
	"TiBO_API/repository/databases/moviesRepo"
	"TiBO_API/repository/databases/usersRepo"
	"time"

	"gorm.io/gorm"
)

type Invoices struct {
	ID               uint                `gorm:"primaryKey"`
	Date             time.Time           `json:"date"`
	ShowTime         string              `json:"show_time"`
	TotalTicket      int                 `json:"total_ticket"`
	TotalTicketPrice int                 `json:"total_ticket_price"`
	TotalPrices      int                 `json:"total_prices"`
	AdminFee         int                 `json:"admin_fee"`
	MovieID          uint                `json:"movie_id"`
	Movie            moviesRepo.Movies   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	UserID           uint                `json:"customer_id"`
	MoviePrice       int                 `json:"movie_price"`
	MovieUrl         string              `json:"movie_url"`
	MovieTitle       string              `json:"movie_title"`
	CinemaAddress    string              `json:"cinema_address"`
	User             usersRepo.Users     `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	CinemaID         uint                `json:"cinema_id"`
	Cinema           cinemasRepo.Cinemas `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CreatedAt        time.Time           `json:"created_at"`
}

func (rec *Invoices) toDomain() invoiceEntity.Domain {
	return invoiceEntity.Domain{
		ID:               rec.ID,
		Date:             rec.Date,
		ShowTime:         rec.ShowTime,
		TotalTicket:      rec.TotalTicket,
		TotalTicketPrice: rec.TotalTicketPrice,
		AdminFee:         rec.AdminFee,
		CreatedAt:        rec.CreatedAt,
		UserID:           rec.UserID,
		CinemaID:         rec.CinemaID,
		CinemaName:       rec.Cinema.Name,
		CinemaAddress:    rec.CinemaAddress,
		MovieID:          rec.MovieID,
		MovieTitle:       rec.MovieTitle,
		MoviePrice:       rec.MoviePrice,
		MovieUrl:         rec.MovieUrl,
		TotalPrices:      rec.TotalPrices,
	}
}

func toDomainArray(rec []Invoices) []invoiceEntity.Domain {
	domain := []invoiceEntity.Domain{}

	for _, val := range rec {
		domain = append(domain, val.toDomain())
	}
	return domain
}

func fromDomain(domain invoiceEntity.Domain) *Invoices {
	return &Invoices{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UserID:    domain.UserID,
		Cinema: cinemasRepo.Cinemas{
			Model: gorm.Model{
				ID: domain.CinemaID,
			},
			Name: domain.CinemaName,
		},
		MovieID: domain.MovieID,
		Movie: moviesRepo.Movies{
			Model: gorm.Model{
				ID: domain.MovieID,
			},
		},
		TotalPrices:      domain.TotalPrices,
		MovieUrl:         domain.MovieUrl,
		MoviePrice:       domain.MoviePrice,
		MovieTitle:       domain.MovieTitle,
		CinemaAddress:    domain.CinemaAddress,
		AdminFee:         domain.AdminFee,
		TotalTicket:      domain.TotalTicket,
		TotalTicketPrice: domain.TotalTicketPrice,
		ShowTime:         domain.ShowTime,
		Date:             domain.Date,
	}
}
