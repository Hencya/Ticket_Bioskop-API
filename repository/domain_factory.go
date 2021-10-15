package repository

import (
	"TiBO_API/businesses/usersEntity"
	"TiBO_API/repository/databases/usersRepo"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) usersEntity.Repository {
	return usersRepo.NewUsersRepository(db)
}
