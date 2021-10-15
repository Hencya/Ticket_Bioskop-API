package cinemasRepo

import (
	"TiBO_API/businesses/cinemasEntity"
	"context"
	"errors"

	"github.com/gosimple/slug"

	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type CinemsRepository struct {
	db *gorm.DB
	//cld *cloudinary.Cloudinary
}

func NewCinemasRepository(db *gorm.DB) cinemasEntity.Repository {
	return &CinemsRepository{
		db: db,
		//cld: cld,
	}
}

func (r *CinemsRepository) PostNewCinema(ctx context.Context, cinemaDomain *cinemasEntity.Domain) (cinemasEntity.Domain, error) {
	rec := FromDomain(*cinemaDomain)
	rec.Slug = slug.Make(rec.Name)

	err := r.db.Create(&rec).Error
	if err != nil {
		return cinemasEntity.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (r *CinemsRepository) GetByAddress(ctx context.Context, addressID []uint) ([]cinemasEntity.Domain, error) {
	rec := []Cinemas{}
	err := r.db.Find(&rec, "addresses_id IN ?", addressID).Error
	if len(rec) == 0 {
		err = errors.New("Not Found")
		return nil, err
	}
	result := ToDomainArray(rec)
	return result, nil
}

func (r *CinemsRepository) GetByName(ctx context.Context, name string) ([]cinemasEntity.Domain, error) {
	rec := []Cinemas{}

	err := r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, err
	}
	result := ToDomainArray(rec)
	return result, nil
}

func (r *CinemsRepository) GetBySlug(ctx context.Context, slug string) (cinemasEntity.Domain, error) {
	rec := Cinemas{}

	err := r.db.Where("slug = ?", slug).First(&rec).Error
	if err != nil {
		return cinemasEntity.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (r *CinemsRepository) Update(ctx context.Context, slugID string, cinemaDomain *cinemasEntity.Domain) (cinemasEntity.Domain, error) {
	rec := FromDomain(*cinemaDomain)
	recData := *rec
	recData.Slug = slug.Make(recData.Name)

	if err := r.db.First(&rec, "slug = ?", slugID).Updates(recData).Error; err != nil {
		return cinemasEntity.Domain{}, err
	}

	return recData.ToDomain(), nil
}

func (r *CinemsRepository) Delete(ctx context.Context, slug string) (string, error) {
	rec := Cinemas{}

	if err := r.db.Delete(&rec, "slug = ?", slug).Error; err != nil {
		return "", err
	}
	return "Cinemas was Deleted", nil
}

func (r *CinemsRepository) GetByID(ctx context.Context, id uint) (cinemasEntity.Domain, error) {
	rec := Cinemas{}

	if err := r.db.Where("ID = ?", id).First(&rec).Error; err != nil {
		return cinemasEntity.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (r *CinemsRepository) FindStatusByTitle(ctx context.Context, name string) (cinemasEntity.Domain, error) {
	rec := Cinemas{}

	if err := r.db.Where("name = ?", name).First(&rec).Error; err != nil {
		return cinemasEntity.Domain{}, err
	}

	return rec.ToDomain(), nil
}
