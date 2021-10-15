package moviesRepo

import (
	moviesEntity "TiBO_API/businesses/movieEntity"
	"TiBO_API/repository/databases/cinemasRepo"

	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	Title            string              `json:"title"`
	Slug             string              `json:"slug"`
	TrailerUrl       string              `json:"trailer_url"`
	MovieUrl         string              `json:"movie_url"`
	Poster           string              `json:"poster"`
	Synopsis         string              `json:"synopsis"`
	Genre            string              `json:"genre"`
	Duration         string              `json:"duration"`
	Language         string              `json:"language"`
	Director         string              `json:"director"`
	CensorRating     string              `json:"censor_rating"`
	Subtitle         string              `json:"subtitle"`
	ScheduleDate     string              `json:"schedule_date"`
	ScheduleTime     string              `json:"schedule_time"`
	StatusComingSoon bool                `json:"status_coming_soon"`
	Ticket           int                 `json:"ticket"`
	Price            int                 `json:"price"`
	CinemasID        uint                `json:"cinema_id"`
	CinemasName      string              `json:"cinemas_name"`
	CinemaAddress    string              `json:"cinema_address"`
	Cinemas          cinemasRepo.Cinemas `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
}

func (rec *Movies) ToDomain() moviesEntity.Domain {
	return moviesEntity.Domain{
		ID:               rec.ID,
		Title:            rec.Title,
		Slug:             rec.Slug,
		TrailerUrl:       rec.TrailerUrl,
		MovieUrl:         rec.MovieUrl,
		Poster:           rec.Poster,
		Synopsis:         rec.Synopsis,
		Genre:            rec.Genre,
		Duration:         rec.Duration,
		Language:         rec.Language,
		Director:         rec.Director,
		CensorRating:     rec.CensorRating,
		Subtitle:         rec.Subtitle,
		ScheduleDate:     rec.ScheduleDate,
		ScheduleTime:     rec.ScheduleTime,
		StatusComingSoon: rec.StatusComingSoon,
		Ticket:           rec.Ticket,
		Price:            rec.Price,
		CinemaID:         rec.CinemasID,
		CinemaName:       rec.CinemasName,
		CinemaAddress:    rec.CinemaAddress,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

func ToDomainArray(rec []Movies) []moviesEntity.Domain {
	domain := []moviesEntity.Domain{}

	for _, val := range rec {
		domain = append(domain, val.ToDomain())
	}
	return domain
}

func FromDomain(domain moviesEntity.Domain) *Movies {
	return &Movies{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Title:            domain.Title,
		Slug:             domain.Slug,
		TrailerUrl:       domain.TrailerUrl,
		Poster:           domain.Poster,
		Synopsis:         domain.Synopsis,
		Genre:            domain.Genre,
		Duration:         domain.Duration,
		Language:         domain.Language,
		Director:         domain.Director,
		CensorRating:     domain.CensorRating,
		Subtitle:         domain.Subtitle,
		ScheduleDate:     domain.ScheduleDate,
		ScheduleTime:     domain.ScheduleTime,
		StatusComingSoon: domain.StatusComingSoon,
		Ticket:           domain.Ticket,
		Price:            domain.Price,
		CinemasID:        domain.CinemaID,
		CinemasName:      domain.CinemaName,
		CinemaAddress:    domain.CinemaAddress,
		MovieUrl:         domain.MovieUrl,
	}
}
