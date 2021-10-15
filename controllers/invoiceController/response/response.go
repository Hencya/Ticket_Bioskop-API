package response

import (
	"TiBO_API/businesses/invoiceEntity"
	"time"
)

type Invoices struct {
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

func FromDomain(domain invoiceEntity.Domain) Invoices {
	return Invoices{
		ID:               domain.ID,
		Date:             domain.Date,
		UserID:           domain.UserID,
		ShowTime:         domain.ShowTime,
		TotalTicket:      domain.TotalTicket,
		TotalTicketPrice: domain.TotalTicketPrice,
		TotalPrices:      domain.TotalPrices,
		AdminFee:         domain.AdminFee,
		MovieID:          domain.MovieID,
		MoviePrice:       domain.MoviePrice,
		MovieTitle:       domain.MovieTitle,
		MovieUrl:         domain.MovieUrl,
		CinemaID:         domain.CinemaID,
		CinemaName:       domain.CinemaName,
		CinemaAddress:    domain.CinemaAddress,
		CreatedAt:        domain.CreatedAt,
	}
}

func FromDomainArray(domain []invoiceEntity.Domain) []Invoices {
	res := []Invoices{}
	for _, val := range domain {
		res = append(res, FromDomain(val))
	}
	return res
}
