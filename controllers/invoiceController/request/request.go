package request

import (
	"TiBO_API/businesses/invoiceEntity"
	"time"
)

type Invoices struct {
	ID               uint      `json:"id"`
	Date             time.Time `json:"date"`
	ShowTime         string    `json:"show_time"`
	TotalTicketPrice int       `json:"total_ticket_price"`
	TotalPrices      int       `json:"total_prices"`
	AdminFee         int       `json:"admin_fee"`
	MoviePrice       int       `json:"movie_price"`
	MovieUrl         string    `json:"movie_url"`
	UserID           uint      `json:"user_id"`
	CinemaName       string    `json:"cinema_name"`
	CinemaAddress    string    `json:"cinema_address"`
	CreatedAt        time.Time `json:"created_at"`
	CinemaID         uint      `json:"cinema_id"`
	MovieID          uint      `json:"movie_id"`
	MovieTitle       string    `json:"movie_title"`
	TotalTicket      int       `json:"total_ticket"`
}

func (req *Invoices) ToDomain() *invoiceEntity.Domain {
	return &invoiceEntity.Domain{
		Date:          req.Date,
		ShowTime:      req.ShowTime,
		TotalPrices:   req.TotalPrices,
		AdminFee:      req.AdminFee,
		MoviePrice:    req.MoviePrice,
		MovieUrl:      req.MovieUrl,
		UserID:        req.UserID,
		CinemaName:    req.CinemaName,
		CinemaAddress: req.CinemaAddress,
		CinemaID:      req.CinemaID,
		MovieID:       req.MovieID,
		MovieTitle:    req.MovieTitle,
		TotalTicket:   req.TotalTicket,
	}
}
