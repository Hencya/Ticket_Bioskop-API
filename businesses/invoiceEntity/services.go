package invoiceEntity

import (
	"TiBO_API/businesses"
	"TiBO_API/businesses/cinemasEntity"
	moviesEntity "TiBO_API/businesses/movieEntity"
	"TiBO_API/businesses/usersEntity"
	"context"
	"errors"
	"time"
)

type InvoiceServices struct {
	invoiceRepository Repository
	moviesRepository  moviesEntity.Repository
	cinemaRepository  cinemasEntity.Repository
	userRepository    usersEntity.Repository
	ContextTimeout    time.Duration
}

func NewInvoiceService(invoiceRepo Repository, movieRepo moviesEntity.Repository,
	cinemaRepo cinemasEntity.Repository, userRepo usersEntity.Repository, timeout time.Duration) Service {
	return &InvoiceServices{
		invoiceRepository: invoiceRepo,
		moviesRepository:  movieRepo,
		cinemaRepository:  cinemaRepo,
		userRepository:    userRepo,
		ContextTimeout:    timeout,
	}
}

func (is *InvoiceServices) Create(ctx context.Context, uuid string, invoiceData *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, is.ContextTimeout)
	defer cancel()

	movieData, err := is.moviesRepository.GetOneByTitle(ctx, invoiceData.MovieTitle)
	if movieData.StatusComingSoon == true {
		return &Domain{}, businesses.ErrMovieNotAvailable
	}

	cinemaOpen, err := is.cinemaRepository.FindStatusByTitle(ctx, movieData.CinemaName)
	if err != nil {
		return &Domain{}, businesses.ErrInternalServer
	}

	if cinemaOpen.IsOpen == false {
		return &Domain{}, businesses.ErrCinemaNotAvailable
	}

	if invoiceData.TotalTicket > movieData.Ticket {
		return &Domain{}, businesses.ErrTicketExceed
	}

	user, err := is.userRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return &Domain{}, errors.New("User is not login")
	}

	//parsing data
	invoiceData.UserID = user.ID
	invoiceData.Date = time.Now()
	invoiceData.ShowTime = movieData.ScheduleTime
	invoiceData.TotalTicketPrice = invoiceData.TotalTicket * movieData.Price
	invoiceData.TotalPrices = invoiceData.TotalTicketPrice + invoiceData.AdminFee
	invoiceData.MoviePrice = movieData.Price
	invoiceData.MovieTitle = movieData.Title
	invoiceData.MovieUrl = movieData.TrailerUrl
	invoiceData.MovieID = movieData.ID
	invoiceData.MovieTitle = movieData.Title
	invoiceData.AdminFee = int(float64(invoiceData.TotalPrices) * 0.05)
	invoiceData.CinemaID = movieData.CinemaID
	invoiceData.CinemaName = movieData.CinemaName
	invoiceData.CinemaAddress = movieData.CinemaAddress

	res, err := is.invoiceRepository.Create(ctx, invoiceData)
	if err != nil {
		return &Domain{}, businesses.ErrInternalServer
	}

	return &res, nil
}

func (is *InvoiceServices) GetByUserID(ctx context.Context, userId string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, is.ContextTimeout)
	defer cancel()

	user, err := is.userRepository.GetByUuid(ctx, userId)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundUser
	}

	res, err := is.invoiceRepository.GetID(ctx, user.ID)
	if err != nil {
		return []Domain{}, businesses.ErrInvoiceNotFound
	}

	return res, nil
}
