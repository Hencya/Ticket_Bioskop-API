package cinemasEntity

import (
	"TiBO_API/businesses"
	"TiBO_API/businesses/addressesEntity"
	"TiBO_API/businesses/geolocationEntity"
	"context"
	"time"
)

type CinemaServices struct {
	CinemaRepository    Repository
	AddressesRepository addressesEntity.Repository
	GeoRepository       geolocationEntity.Repository
	ContextTimeout      time.Duration
}

func NewCinemaServices(repoCinema Repository, repoAddresses addressesEntity.Repository,
	geoRepo geolocationEntity.Repository, timeout time.Duration) Service {
	return &CinemaServices{
		CinemaRepository:    repoCinema,
		AddressesRepository: repoAddresses,
		GeoRepository:       geoRepo,
		ContextTimeout:      timeout,
	}
}

func (cs *CinemaServices) CreateCinema(ctx context.Context, cinemaData *Domain, addressData *addressesEntity.Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	newAddr, err := cs.AddressesRepository.Insert(addressData)
	cinemaData.AddressID = newAddr.ID

	res, err := cs.CinemaRepository.PostNewCinema(ctx, cinemaData)
	if res == (Domain{}) {
		return &Domain{}, businesses.ErrDuplicateData
	}
	if err != nil {
		return &Domain{}, err
	}
	return &res, nil
}

func (cs *CinemaServices) FindByIP(ctx context.Context, ip string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	location, err := cs.GeoRepository.GetLocationByIP(ctx, ip)
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	addrData, err := cs.AddressesRepository.FindByCity(location.City)
	if err != nil {
		return []Domain{}, businesses.ErrNearestCinemaNotFound
	}

	addressID := []uint{}
	for _, val := range addrData {
		addressID = append(addressID, val.ID)
	}
	res, err := cs.CinemaRepository.GetByAddress(ctx, addressID)

	if err != nil {
		return []Domain{}, businesses.ErrNearestCinemaNotFound
	}
	return res, nil

}

func (cs *CinemaServices) FindByName(ctx context.Context, name string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	res, err := cs.CinemaRepository.GetByName(ctx, name)
	if err != nil {
		return []Domain{}, businesses.ErrCinemaNotFound
	}
	return res, nil
}

func (cs *CinemaServices) FindBySlug(ctx context.Context, slug string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	res, err := cs.CinemaRepository.GetBySlug(ctx, slug)
	if err != nil {
		return Domain{}, businesses.ErrCinemaNotFound
	}
	return res, nil
}

func (cs *CinemaServices) UpdateCinema(ctx context.Context, cinemaData *Domain, addressData *addressesEntity.Domain, slug string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	newAddr, err := cs.AddressesRepository.Insert(addressData)
	cinemaData.AddressID = newAddr.ID

	res, err := cs.CinemaRepository.Update(ctx, slug, cinemaData)
	if err != nil {
		return &Domain{}, businesses.ErrCinemaNotFound
	}
	return &res, nil
}

func (cs *CinemaServices) DeleteBySlug(ctx context.Context, slug string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	res, err := cs.CinemaRepository.Delete(ctx, slug)
	if err != nil {
		return "", businesses.ErrCinemaNotFound
	}
	return res, nil
}
