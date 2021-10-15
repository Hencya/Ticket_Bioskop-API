package geolocation

import (
	"TiBO_API/businesses/geolocationEntity"
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

type GeolocationRepository struct {
	httpClient http.Client
}

func NewGeolocationRepository() geolocationEntity.Repository {
	return &GeolocationRepository{
		httpClient: http.Client{},
	}
}

func (geo *GeolocationRepository) GetLocationByIP(ctx context.Context, ip string) (geolocationEntity.Domain, error) {

	//req, _ := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	//req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	//resp, err := geo.httpClient.Do(req)
	//if err != nil {
	//	return geolocationEntity.Domain{}, err
	//}
	//
	//defer resp.Body.Close()
	//
	//data := Response{}
	//err = json.NewDecoder(resp.Body).Decode(&data)
	//if err != nil {
	//	return geolocationEntity.Domain{}, err
	//}
	//return data.toDomain(), nil
	req, _ := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := geo.httpClient.Do(req)
	if err != nil {
		return geolocationEntity.Domain{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return geolocationEntity.Domain{}, err
	}
	return data.toDomain(), nil
}
