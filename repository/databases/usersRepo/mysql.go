package usersRepo

import (
	"TiBO_API/businesses/usersEntity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) usersEntity.Repository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) CreateNewUser(ctx context.Context, usersDomain *usersEntity.Domain) (*usersEntity.Domain, error) {
	rec := FromDomain(usersDomain)
	rec.Uuid, _ = uuid.NewRandom()
	rec.Role = "user"

	err := r.db.Create(&rec).Error
	if err != nil {
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (usersEntity.Domain, error) {
	rec := Users{}

	err := r.db.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return usersEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *UsersRepository) GetByUuid(ctx context.Context, uuid string) (usersEntity.Domain, error) {
	rec := Users{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return usersEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *UsersRepository) UpdateUser(ctx context.Context, id string, usersDomain *usersEntity.Domain) (*usersEntity.Domain, error) {
	rec := FromDomain(usersDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &usersEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &usersEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *UsersRepository) UploadAvatar(ctx context.Context, id string, usersDomain *usersEntity.Domain) (*usersEntity.Domain, error) {
	rec := FromDomain(usersDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &usersEntity.Domain{}, err
	}

	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &usersEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil
}

func (r *UsersRepository) DeleteUserByUuid(ctx context.Context, id string) (string, error) {
	rec := Users{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "User was Deleted", nil
}
