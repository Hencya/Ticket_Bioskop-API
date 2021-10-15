package request

import moviesEntity "TiBO_API/businesses/movieEntity"

type Movies struct {
	Slug             string `json:"slug"`
	Title            string `json:"title"`
	TrailerUrl       string `json:"trailer_url"`
	Poster           string `json:"poster"`
	Synopsis         string `json:"synopsis"`
	Genre            string `json:"genre"`
	Duration         string `json:"duration"`
	Language         string `json:"language"`
	Director         string `json:"director"`
	CensorRating     string `json:"censor_rating"`
	Subtitle         string `json:"subtitle"`
	ScheduleDate     string `json:"schedule_date"`
	ScheduleTime     string `json:"schedule_time"`
	StatusComingSoon bool   `json:"status_coming_soon"`
	AdminFee         int    `json:"admin_fee"`
	Ticket           int    `json:"ticket"`
	Price            int    `json:"price"`
	MovieUrl         string `json:"movie_url"`
}

func (req *Movies) ToDomain() *moviesEntity.Domain {
	return &moviesEntity.Domain{
		Title:            req.Title,
		Slug:             req.Slug,
		TrailerUrl:       req.TrailerUrl,
		Poster:           req.Poster,
		Synopsis:         req.Synopsis,
		Genre:            req.Genre,
		Duration:         req.Duration,
		Language:         req.Language,
		Director:         req.Director,
		CensorRating:     req.CensorRating,
		Subtitle:         req.Subtitle,
		ScheduleDate:     req.ScheduleDate,
		ScheduleTime:     req.ScheduleTime,
		StatusComingSoon: req.StatusComingSoon,
		Ticket:           req.Ticket,
		Price:            req.Price,
		MovieUrl:         req.MovieUrl,
	}
}
