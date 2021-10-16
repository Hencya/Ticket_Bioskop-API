package moviesEntity_test

import (
mockCinema "TiBO_API/businesses/CinemasEntity/mocks"
"TiBO_API/businesses/addressesEntity"
mockAddress "TiBO_API/businesses/addressesEntity/mocks"
"TiBO_API/businesses/cinemasEntity"
moviesEntity "TiBO_API/businesses/movieEntity"
mockMovie "TiBO_API/businesses/movieEntity/mocks"
	"TiBO_API/businesses/usersEntity"
	"context"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
"testing"
"time"
)

var (
	mockMovieRepository mockMovie.Repository
	mockCinemaRepository mockCinema.Repository
	mockAddressesepository mockAddress.Repository
	movieService moviesEntity.Service
	movieDomain moviesEntity.Domain
	cinemaDomain cinemasEntity.Domain
	movieTrueDomain moviesEntity.Domain
	movieNewDomain moviesEntity.Domain
)

func TestMain(m *testing.M) {
	movieService = moviesEntity.NewMoviesServices(&mockMovieRepository,&mockCinemaRepository,&mockAddressesepository,
		time.Second*2)
	movieDomain = moviesEntity.Domain{
		ID:               1,
		Slug:             "test-slug",
		Title:            "testing title",
		TrailerUrl:       "www.testingTrailer.com",
		MovieUrl:         "www.testingMovie.com",
		Poster:           "images/poster/poster.png",
		Synopsis:         "ini adalah testing",
		Genre:            "Horror",
		Duration:         "3 Hours",
		Language:         "English",
		Director:         "Faiz",
		CensorRating:     "20+",
		Subtitle:         "Arabic",
		ScheduleDate:     "19-Januari-2021",
		ScheduleTime:     "18:00",
		StatusComingSoon: false,
		Ticket:           100,
		Price:            10000,
		CinemaID:         1,
		CinemaName:       "testing cinema",
		CinemaAddress:    "jl testing",
		AdminFee:         1000,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	movieTrueDomain = moviesEntity.Domain{
		ID:               1,
		Slug:             "test-slug",
		Title:            "testing title",
		TrailerUrl:       "www.testingTrailer.com",
		MovieUrl:         "www.testingMovie.com",
		Poster:           "poster.jpg",
		Synopsis:         "ini adalah testing",
		Genre:            "Horror",
		Duration:         "3 Hours",
		Language:         "English",
		Director:         "Faiz",
		CensorRating:     "20+",
		Subtitle:         "Arabic",
		ScheduleDate:     "19-Januari-2021",
		ScheduleTime:     "18:00",
		StatusComingSoon: true,
		Ticket:           100,
		Price:            10000,
		CinemaID:         1,
		CinemaName:       "testing cinema",
		CinemaAddress:    "jl testing",
		AdminFee:         1000,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	movieNewDomain = moviesEntity.Domain{
		ID:               1,
		Slug:             "test-slug",
		Title:            "testing title",
		TrailerUrl:       "www.testingTrailer.com",
		MovieUrl:         "www.testingMovie.com",
		Poster:           "poster.jpg",
		Synopsis:         "ini adalah testing",
		Genre:            "Horror",
		Duration:         "3 Hours",
		Language:         "English",
		Director:         "Faiz",
		CensorRating:     "20+",
		Subtitle:         "Arabic",
		ScheduleDate:     "19-Januari-2021",
		ScheduleTime:     "18:00",
		StatusComingSoon: true,
		Ticket:           100,
		Price:            10000,
		CinemaID:         1,
		CinemaName:       "testing cinema",
		CinemaAddress:    "jl testing",
		AdminFee:         1000,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	cinemaDomain = cinemasEntity.Domain{
		ID:          1,
		Slug:        "asd;sasjdlkajd",
		Name:        "Test Cinema",
		PhoneNumber: "(021)123123",
		IsOpen:      true,
		AddressID:   2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m.Run()
}

func TestMoviesServices_CreateMovie(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockCinemaRepository.On("GetBySlug",mock.Anything, mock.AnythingOfType("string")).Return(cinemasEntity.Domain{}, nil).Once()
		mockAddressesepository.On("FindByID", mock.AnythingOfType("uint")).Return(addressesEntity.Domain{}, nil).Once()
		mockMovieRepository.On("PostNewMovie", mock.Anything,mock.Anything).Return(movieDomain, nil).Once()

		input := &moviesEntity.Domain{
			Slug:             "test-slug",
			Title:            "testing title",
			TrailerUrl:       "www.testingTrailer.com",
			MovieUrl:         "www.testingMovie.com",
			Poster:           "poster.jpg",
			Synopsis:         "ini adalah testing",
			Genre:            "Horror",
			Duration:         "3 Hours",
			Language:         "English",
			Director:         "Faiz",
			CensorRating:     "20+",
			Subtitle:         "Arabic",
			ScheduleDate:     "19-Januari-2021",
			ScheduleTime:     "18:00",
			StatusComingSoon: true,
			Ticket:           100,
			Price:            10000,
			CinemaID:         1,
			CinemaName:       "testing cinema",
			CinemaAddress:    "jl testing",
			AdminFee:         1000,
		}
		resp, err := movieService.CreateMovie(context.Background(),input,"test-slug")

		assert.Nil(t, err)
		assert.Equal(t, &movieDomain, resp)
	})
	t.Run("Invalid Test | slug ", func(t *testing.T){
		mockCinemaRepository.On("GetBySlug",mock.Anything, mock.AnythingOfType("string")).Return(cinemasEntity.Domain{}, assert.AnError).Once()
		mockAddressesepository.On("FindByID", mock.AnythingOfType("uint")).Return(addressesEntity.Domain{}, nil).Once()
		mockMovieRepository.On("PostNewMovie", mock.Anything,mock.Anything).Return(moviesEntity.Domain{}, nil).Once()

		input := &moviesEntity.Domain{
			Slug:             "test-slug",
			Title:            "testing title",
			TrailerUrl:       "www.testingTrailer.com",
			MovieUrl:         "www.testingMovie.com",
			Poster:           "poster.jpg",
			Synopsis:         "ini adalah testing",
			Genre:            "Horror",
			Duration:         "3 Hours",
			Language:         "English",
			Director:         "Faiz",
			CensorRating:     "20+",
			Subtitle:         "Arabic",
			ScheduleDate:     "19-Januari-2021",
			ScheduleTime:     "18:00",
			StatusComingSoon: true,
			Ticket:           100,
			Price:            10000,
			CinemaID:         1,
			CinemaName:       "testing cinema",
			CinemaAddress:    "jl testing",
			AdminFee:         1000,
		}

		resp, err := movieService.CreateMovie(context.Background(),input,"test-slug")

		assert.NotNil(t, err)
		assert.NotEqual(t, movieDomain, resp)
	})
	t.Run("Invalid Test | ID ", func(t *testing.T){
		mockCinemaRepository.On("GetBySlug",mock.Anything, mock.AnythingOfType("string")).Return(cinemasEntity.Domain{}, assert.AnError).Once()
		mockAddressesepository.On("FindByID", mock.AnythingOfType("uint")).Return(addressesEntity.Domain{}, assert.AnError).Once()
		mockMovieRepository.On("PostNewMovie", mock.Anything,mock.Anything).Return(moviesEntity.Domain{}, nil).Once()

		input := &moviesEntity.Domain{
			Slug:             "test-slug",
			Title:            "testing title",
			TrailerUrl:       "www.testingTrailer.com",
			MovieUrl:         "www.testingMovie.com",
			Poster:           "poster.jpg",
			Synopsis:         "ini adalah testing",
			Genre:            "Horror",
			Duration:         "3 Hours",
			Language:         "English",
			Director:         "Faiz",
			CensorRating:     "20+",
			Subtitle:         "Arabic",
			ScheduleDate:     "19-Januari-2021",
			ScheduleTime:     "18:00",
			StatusComingSoon: true,
			Ticket:           100,
			Price:            10000,
			CinemaID:         1,
			CinemaName:       "testing cinema",
			CinemaAddress:    "jl testing",
			AdminFee:         1000,
		}

		resp, err := movieService.CreateMovie(context.Background(),input,"test-slug")

		assert.NotNil(t, err)
		assert.NotEqual(t, movieDomain, resp)
	})
}

func TestMoviesServices_FindByTitle(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockMovieRepository.On("GetByTitle",  mock.Anything,mock.AnythingOfType("string")).Return([]moviesEntity.Domain{movieDomain}, nil).Once()

		resp, err := movieService.FindByTitle(context.Background(),"testing title")

		assert.Nil(t, err)
		assert.Contains(t, resp, movieDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockMovieRepository.On("GetByTitle",  mock.Anything,mock.AnythingOfType("string")).Return([]moviesEntity.Domain{}, assert.AnError).Once()

		resp, err := movieService.FindByTitle(context.Background(),"testing title")

		assert.NotNil(t, err)
		assert.Equal(t, []moviesEntity.Domain{}, resp)
	})
}

func TestMoviesServices_GetOneByTitle(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockMovieRepository.On("GetOneByTitle",  mock.Anything,mock.Anything).Return(movieDomain, nil).Once()

		resp, err := movieService.GetOneByTitle(context.Background(),"testing title")

		assert.Nil(t, err)
		assert.Equal(t, resp, movieDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockMovieRepository.On("GetOneByTitle",  mock.Anything,mock.Anything).Return(moviesEntity.Domain{}, assert.AnError).Once()

		resp, err := movieService.GetOneByTitle(context.Background(),"testing title")

		assert.NotNil(t, err)
		assert.Equal(t, moviesEntity.Domain{}, resp)
	})
}

func TestMoviesServices_FindBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockMovieRepository.On("GetBySlug",  mock.Anything,mock.AnythingOfType("string")).Return(movieDomain, nil).Once()

		resp, err := movieService.FindBySlug(context.Background(),"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, resp, movieDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockMovieRepository.On("GetBySlug",  mock.Anything,mock.AnythingOfType("string")).Return(moviesEntity.Domain{}, assert.AnError).Once()

		resp, err := movieService.FindBySlug(context.Background(),"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, moviesEntity.Domain{}, resp)
	})
}

func TestMoviesServices_DeleteBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockMovieRepository.On("Delete", mock.Anything,mock.AnythingOfType("string")).Return("Movie was Deleted", nil).Once()

		resp, err := movieService.DeleteBySlug(context.Background(),"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, "Movie was Deleted", resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockMovieRepository.On("Delete", mock.Anything,mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		resp, err := movieService.DeleteBySlug(context.Background(),"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockCinemaRepository.On("GetBySlug",mock.Anything, mock.AnythingOfType("string")).Return(cinemasEntity.Domain{}, nil).Once()
		mockAddressesepository.On("FindByID", mock.AnythingOfType("uint")).Return(addressesEntity.Domain{}, nil).Once()
		mockMovieRepository.On("Update", mock.Anything,mock.Anything,mock.Anything).Return(movieDomain, nil).Once()

		input := &moviesEntity.Domain{
			Slug:             "test-slug",
			Title:            "testing title",
			TrailerUrl:       "www.testingTrailer.com",
			MovieUrl:         "www.testingMovie.com",
			Poster:           "poster.jpg",
			Synopsis:         "ini adalah testing",
			Genre:            "Horror",
			Duration:         "3 Hours",
			Language:         "English",
			Director:         "Faiz",
			CensorRating:     "20+",
			Subtitle:         "Arabic",
			ScheduleDate:     "19-Januari-2021",
			ScheduleTime:     "18:00",
			StatusComingSoon: true,
			Ticket:           100,
			Price:            10000,
			CinemaID:         1,
			CinemaName:       "testing cinema",
			CinemaAddress:    "jl testing",
			AdminFee:         1000,
		}
		resp, err := movieService.UpdateMovie(context.Background(),input,"test-slug","asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, &movieDomain, resp)
	})
}

func TestMoviesServices_UploadPoster(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockMovieRepository.On("GetBySlug", mock.Anything,mock.AnythingOfType("string")).Return(movieDomain, nil).Once()
		mockMovieRepository.On("UploadPoster",mock.Anything,mock.Anything,mock.Anything).Return(&movieDomain, nil).Once()

		req := &moviesEntity.Domain{
			Poster: "images/poster/poster.png",
		}

		res, err := movieService.UploadPoster(context.Background(),"test-slug",req.Poster)

		assert.Nil(t, err)
		assert.Equal(t, movieDomain.Poster, res.Poster)
	})
	t.Run("Invalid Test || movie not found", func(t *testing.T){
		mockMovieRepository.On("GetBySlug", mock.Anything,mock.AnythingOfType("string")).Return(moviesEntity.Domain{}, assert.AnError).Once()
		mockMovieRepository.On("UploadPoster",mock.Anything,mock.Anything,mock.Anything).Return(&movieDomain, nil).Once()

		req := &moviesEntity.Domain{
			Poster: "images/poster/poster.png",
		}

		res, err := movieService.UploadPoster(context.Background(),"test-slug",req.Poster)

		assert.NotNil(t, err)
		assert.NotEqual(t, usersEntity.Domain{}, res)
	})
	t.Run("Invalid Test || failed to upload poster", func(t *testing.T){
		mockMovieRepository.On("GetBySlug", mock.Anything,mock.AnythingOfType("string")).Return(movieDomain,nil).Once()
		mockMovieRepository.On("UploadPoster",mock.Anything,mock.Anything,mock.Anything).Return(&moviesEntity.Domain{}, assert.AnError).Once()

		req := &moviesEntity.Domain{
			Poster: "images/poster/poster.png",
		}

		res, _ := movieService.UploadPoster(context.Background(),"test-slug",req.Poster)

		assert.NotEqual(t, usersEntity.Domain{}, res)
	})
}