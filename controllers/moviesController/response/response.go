package response

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
	Ticket           int    `json:"ticket"`
	Price            int    `json:"price"`
	AdminFee         int    `json:"admin_fee"`
	CinemaID         uint   `json:"cinema_id"`
	CinemaName       string `json:"cinema_name"`
	CinemaAddress    string `json:"cinema_address"`
}

func FromDomain(domain moviesEntity.Domain) Movies {
	return Movies{
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
		CinemaID:         domain.CinemaID,
		CinemaName:       domain.CinemaName,
		CinemaAddress:    domain.CinemaAddress,
	}
}

func FromDomainArray(domain []moviesEntity.Domain) []Movies {
	res := []Movies{}
	for _, val := range domain {
		res = append(res, FromDomain(val))
	}
	return res
}
