package request

import (
	"TiBO_API/businesses/addressesEntity"
	"TiBO_API/businesses/cinemasEntity"
	"TiBO_API/controllers/usersController/addressesController/request"
)

type Cinemas struct {
	Slug        string            `json:"slug"`
	Name        string            `json:"name"`
	PhoneNumber string            `json:"phone_number"`
	IsOpen      bool              `json:"is_open"`
	AddressID   uint              `json:"address_id"`
	Addresses   request.Addresses `json:"addresses"`
}

func (req *Cinemas) ToDomain() (*cinemasEntity.Domain, *addressesEntity.Domain) {
	return &cinemasEntity.Domain{
			Slug:        req.Slug,
			Name:        req.Name,
			IsOpen:      req.IsOpen,
			PhoneNumber: req.PhoneNumber,
			AddressID:   req.AddressID,
		}, &addressesEntity.Domain{
			Street:   req.Addresses.Street,
			City:     req.Addresses.City,
			Province: req.Addresses.Province,
		}
}
