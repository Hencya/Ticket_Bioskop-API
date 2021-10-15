package cinemasRepo

import (
	"TiBO_API/businesses/cinemasEntity"
	"TiBO_API/repository/databases/addressesRepo"

	"gorm.io/gorm"
)

type Cinemas struct {
	gorm.Model
	Name        string                  `json:"name"`
	Slug        string                  `json:"slug"`
	IsOpen      bool                    `json:"is_open"`
	PhoneNumber string                  `json:"phone_number"`
	AddressesID uint                    `json:"address_id"`
	Addresses   addressesRepo.Addresses `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Status      bool                    `json:"status"`
}

func (rec *Cinemas) ToDomain() cinemasEntity.Domain {
	return cinemasEntity.Domain{
		ID:          rec.ID,
		Slug:        rec.Slug,
		Name:        rec.Name,
		IsOpen:      rec.IsOpen,
		PhoneNumber: rec.PhoneNumber,
		AddressID:   rec.AddressesID,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func ToDomainArray(rec []Cinemas) []cinemasEntity.Domain {
	domain := []cinemasEntity.Domain{}

	for _, val := range rec {
		domain = append(domain, val.ToDomain())
	}
	return domain
}

func FromDomain(domain cinemasEntity.Domain) *Cinemas {
	return &Cinemas{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Name:        domain.Name,
		Slug:        domain.Slug,
		IsOpen:      domain.IsOpen,
		PhoneNumber: domain.PhoneNumber,
		AddressesID: domain.AddressID,
	}
}
