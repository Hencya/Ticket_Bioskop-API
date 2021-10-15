package moviesRepo

import (
	moviesEntity "TiBO_API/businesses/movieEntity"
	"context"
	"errors"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type MoviesRepository struct {
	db *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) moviesEntity.Repository {
	return &MoviesRepository{
		db: db,
	}
}

func (r *MoviesRepository) PostNewMovie(ctx context.Context, moviesDomain *moviesEntity.Domain) (moviesEntity.Domain, error) {
	rec := FromDomain(*moviesDomain)
	rec.Slug = slug.Make(rec.Title)

	err := r.db.Create(&rec).Error
	if err != nil {
		return moviesEntity.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (r *MoviesRepository) GetByCinemaId(ctx context.Context, cinemaID uint) ([]moviesEntity.Domain, error) {
	rec := []Movies{}

	err := r.db.Find(&rec, "cinema_id IN ?", cinemaID).Error
	if len(rec) == 0 {
		err = errors.New("Not Found")
		return nil, err
	}
	result := ToDomainArray(rec)

	return result, nil
}

func (r *MoviesRepository) GetByTitle(ctx context.Context, title string) ([]moviesEntity.Domain, error) {
	rec := []Movies{}

	err := r.db.Find(&rec, "title LIKE ?", "%"+title+"%").Error
	if err != nil {
		return nil, err
	}
	result := ToDomainArray(rec)

	return result, nil
}

func (r *MoviesRepository) GetOneByTitle(ctx context.Context, title string) (moviesEntity.Domain, error) {
	rec := Movies{}

	err := r.db.Where("title = ?", title).First(&rec).Error
	if err != nil {
		return moviesEntity.Domain{}, err
	}

	result := rec.ToDomain()

	return result, nil
}

func (r *MoviesRepository) GetBySlug(ctx context.Context, slug string) (moviesEntity.Domain, error) {
	rec := Movies{}

	err := r.db.Where("slug = ?", slug).First(&rec).Error
	if err != nil {
		return moviesEntity.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (r *MoviesRepository) Update(ctx context.Context, slugID string, moviesDomain *moviesEntity.Domain) (moviesEntity.Domain, error) {
	rec := FromDomain(*moviesDomain)
	recData := *rec
	recData.Slug = slug.Make(recData.Title)

	if err := r.db.First(&rec, "slug = ?", slugID).Updates(recData).Error; err != nil {
		return moviesEntity.Domain{}, err
	}

	return recData.ToDomain(), nil
}

func (r *MoviesRepository) Delete(ctx context.Context, slug string) (string, error) {
	rec := Movies{}

	if err := r.db.Delete(&rec, "slug = ?", slug).Error; err != nil {
		return "", err
	}
	return "Movies was Deleted", nil
}

func (r *MoviesRepository) UploadPoster(ctx context.Context, slug string, moviesDomain *moviesEntity.Domain) (*moviesEntity.Domain, error) {
	rec := FromDomain(*moviesDomain)

	if err := r.db.Where("slug = ?", slug).Updates(&rec).Error; err != nil {
		return &moviesEntity.Domain{}, err
	}

	if err := r.db.Where("slug = ?", slug).First(&rec).Error; err != nil {
		return &moviesEntity.Domain{}, err
	}

	result := rec.ToDomain()

	return &result, nil
}

func (r *MoviesRepository) GetByMovieId(ctx context.Context, id uint) (moviesEntity.Domain, error) {
	rec := Movies{}

	err := r.db.Where("id = ?", id).Error
	if err != nil {
		return moviesEntity.Domain{}, err
	}
	result := rec.ToDomain()
	return result, nil
}
