package repository

import (
	"TiBO_API/businesses/addressesEntity"
	"TiBO_API/businesses/cinemasEntity"
	"TiBO_API/businesses/geolocationEntity"
	"TiBO_API/businesses/invoiceEntity"
	"TiBO_API/businesses/usersEntity"
	"TiBO_API/repository/databases/addressesRepo"
	"TiBO_API/repository/databases/cinemasRepo"
	"TiBO_API/repository/databases/invoiceRepo"
	"TiBO_API/repository/databases/moviesRepo"
	"TiBO_API/repository/databases/usersRepo"

	moviesEntity "TiBO_API/businesses/movieEntity"
	geolocation "TiBO_API/repository/thirdparties/ipapi"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) usersEntity.Repository {
	return usersRepo.NewUsersRepository(db)
}

func NewCinemasRepository(db *gorm.DB) cinemasEntity.Repository {
	return cinemasRepo.NewCinemasRepository(db)
}

func NewAddressesRepository(db *gorm.DB) addressesEntity.Repository {
	return addressesRepo.NewAddressesRepository(db)
}

func NewGeolocationRepository() geolocationEntity.Repository {
	return geolocation.NewGeolocationRepository()
}

func NewMoviesRepository(db *gorm.DB) moviesEntity.Repository {
	return moviesRepo.NewMoviesRepository(db)
}

func NewInvoicesRepository(db *gorm.DB) invoiceEntity.Repository {
	return invoiceRepo.NewInvoicesRepository(db)
}
