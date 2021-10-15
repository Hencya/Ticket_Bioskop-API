package response

import "TiBO_API/businesses/cinemasEntity"

type Cinemas struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	IsOpen      bool   `json:"is_open"`
	PhoneNumber string `json:"phone_number"`
	AddressID   uint   `json:"addresses"`
}

func FromDomain(domain cinemasEntity.Domain) Cinemas {
	return Cinemas{
		Name:        domain.Name,
		Slug:        domain.Slug,
		IsOpen:      domain.IsOpen,
		PhoneNumber: domain.PhoneNumber,
		AddressID:   domain.AddressID,
	}
}
