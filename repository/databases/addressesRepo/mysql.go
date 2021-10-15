package addressesRepo

import (
	"TiBO_API/businesses/addressesEntity"

	"gorm.io/gorm"
)

type AddressesRepository struct {
	db *gorm.DB
}

func NewAddressesRepository(db *gorm.DB) addressesEntity.Repository {
	return &AddressesRepository{
		db: db,
	}
}

func (mysqlRepo *AddressesRepository) Insert(address *addressesEntity.Domain) (addressesEntity.Domain, error) {
	rec := FromDomain(*address)
	queryString := "street = ? AND city = ? AND province = ?"
	err := mysqlRepo.db.First(&rec, queryString, address.Street, address.City, address.Province).Error
	if err != nil {
		if errCreate := mysqlRepo.db.Create(&rec).Error; err != nil {
			return addressesEntity.Domain{}, errCreate
		}
	}
	return rec.ToDomain(), nil
}

func (mysqlRepo *AddressesRepository) FindByID(id uint) (addressesEntity.Domain, error) {
	rec := Addresses{}
	err := mysqlRepo.db.First(&rec, id).Error
	if err != nil {
		return addressesEntity.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (mysqlRepo *AddressesRepository) FindByCity(city string) ([]addressesEntity.Domain, error) {
	rec := []Addresses{}
	err := mysqlRepo.db.Find(&rec, "city = ?", city).Error
	if err != nil {
		return []addressesEntity.Domain{}, err
	}
	domainAddresses := []addressesEntity.Domain{}
	for _, val := range rec {
		domainAddresses = append(domainAddresses, val.ToDomain())
	}
	return domainAddresses, nil
}
