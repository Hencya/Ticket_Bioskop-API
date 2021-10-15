package addressesRepo

import "TiBO_API/businesses/addressesEntity"

type Addresses struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
}

func (rec *Addresses) ToDomain() addressesEntity.Domain {
	return addressesEntity.Domain{
		ID:       rec.ID,
		Street:   rec.Street,
		City:     rec.City,
		Province: rec.Province,
	}
}

func FromDomain(domain addressesEntity.Domain) *Addresses {
	return &Addresses{
		ID:       domain.ID,
		Street:   domain.Street,
		City:     domain.City,
		Province: domain.Province,
	}
}
