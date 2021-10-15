package moviesEntity

import (
	"context"
	"time"
)

type Domain struct {
	ID               uint
	Slug             string
	Title            string
	TrailerUrl       string
	MovieUrl         string
	Poster           string
	Synopsis         string
	Genre            string
	Duration         string
	Language         string
	Director         string
	CensorRating     string
	Subtitle         string
	ScheduleDate     string
	ScheduleTime     string
	StatusComingSoon bool
	Ticket           int
	Price            int
	CinemaID         uint
	CinemaName       string
	CinemaAddress    string
	AdminFee         int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	CreateMovie(ctx context.Context, movieDomain *Domain, slug string) (*Domain, error)
	FindByTitle(ctx context.Context, title string) ([]Domain, error)
	FindBySlug(ctx context.Context, slug string) (Domain, error)
	UpdateMovie(ctx context.Context, movieDomain *Domain, slugCinema string, slugMovie string) (*Domain, error)
	UploadPoster(ctx context.Context, slug string, fileLocation string) (*Domain, error)
	DeleteBySlug(ctx context.Context, slug string) (string, error)
}

type Repository interface {
	// Database mysql
	PostNewMovie(ctx context.Context, movieDomain *Domain) (Domain, error)
	UploadPoster(ctx context.Context, slug string, data *Domain) (*Domain, error)
	GetByTitle(ctx context.Context, title string) ([]Domain, error)
	GetOneByTitle(ctx context.Context, title string) (Domain, error)
	GetBySlug(ctx context.Context, slug string) (Domain, error)
	Update(ctx context.Context, slug string, movieDomain *Domain) (Domain, error)
	Delete(ctx context.Context, slug string) (string, error)
	GetByMovieId(ctx context.Context, id uint) (Domain, error)
}
