package moviesEntity

import (
	"TiBO_API/businesses"
	"TiBO_API/businesses/addressesEntity"
	"TiBO_API/businesses/cinemasEntity"
	"context"
	"time"
)

type MoviesServices struct {
	MovieRepository   Repository
	CinemaRepository  cinemasEntity.Repository
	AddressRepository addressesEntity.Repository
	ContextTimeout    time.Duration
}

func NewMoviesServices(repoMovie Repository, repoCinema cinemasEntity.Repository, repoAddr addressesEntity.Repository, timeout time.Duration) Service {
	return &MoviesServices{
		MovieRepository:   repoMovie,
		CinemaRepository:  repoCinema,
		AddressRepository: repoAddr,
		ContextTimeout:    timeout,
	}
}

func (ms *MoviesServices) CreateMovie(ctx context.Context, movieDomain *Domain, slug string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	cinema, err := ms.CinemaRepository.GetBySlug(ctx, slug)
	if err != nil {
		return &Domain{}, businesses.ErrCinemaNotFound
	}

	cinemaAddr, err := ms.AddressRepository.FindByID(cinema.AddressID)
	{
		if err != nil {
			return &Domain{}, businesses.ErrCinemaNotFound
		}
	}

	movieDomain.CinemaID = cinema.ID
	movieDomain.CinemaName = cinema.Name
	movieDomain.CinemaAddress = cinemaAddr.Street + cinemaAddr.City + cinemaAddr.Province

	res, err := ms.MovieRepository.PostNewMovie(ctx, movieDomain)
	if res == (Domain{}) {
		return &Domain{}, businesses.ErrInternalServer
	}

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (ms *MoviesServices) FindByTitle(ctx context.Context, title string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	res, err := ms.MovieRepository.GetByTitle(ctx, title)
	if err != nil {
		return []Domain{}, businesses.ErrMovieNotFound
	}

	return res, nil
}

func (ms *MoviesServices) GetOneByTitle(ctx context.Context, title string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	res, err := ms.MovieRepository.GetOneByTitle(ctx, title)
	if err != nil {
		return Domain{}, businesses.ErrMovieNotFound
	}

	return res, nil
}

func (ms *MoviesServices) FindBySlug(ctx context.Context, slug string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	res, err := ms.MovieRepository.GetBySlug(ctx, slug)
	if err != nil {
		return Domain{}, businesses.ErrMovieNotFound
	}

	return res, nil
}

func (ms *MoviesServices) UpdateMovie(ctx context.Context, movieDomain *Domain, slugCinema string, slugMovie string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	cinema, err := ms.CinemaRepository.GetBySlug(ctx, slugCinema)
	if err != nil {
		return &Domain{}, businesses.ErrCinemaNotFound
	}

	cinemaAddr, err := ms.AddressRepository.FindByID(cinema.AddressID)
	{
		if err != nil {
			return &Domain{}, businesses.ErrCinemaNotFound
		}
	}

	movieDomain.CinemaID = cinema.ID
	movieDomain.CinemaName = cinema.Name
	movieDomain.CinemaAddress = cinemaAddr.Street + cinemaAddr.City + cinemaAddr.Province

	res, err := ms.MovieRepository.Update(ctx, slugMovie, movieDomain)
	if res == (Domain{}) {
		return &Domain{}, businesses.ErrDuplicateData
	}

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (ms *MoviesServices) DeleteBySlug(ctx context.Context, slug string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	res, err := ms.MovieRepository.Delete(ctx, slug)
	if err != nil {
		return "", businesses.ErrMovieNotFound
	}
	return res, nil
}

func (ms *MoviesServices) UploadPoster(ctx context.Context, slug string, fileLocation string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ms.ContextTimeout)
	defer cancel()

	movie, err := ms.MovieRepository.GetBySlug(ctx, slug)
	if err != nil {
		return &Domain{}, err
	}

	movie.Poster = fileLocation
	updateAvatar, err := ms.MovieRepository.UploadPoster(ctx, slug, &movie)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}
