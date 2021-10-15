package geolocationEntity

import "context"

type Domain struct {
	IP   string
	City string
}

type Repository interface {
	GetLocationByIP(ctx context.Context, ip string) (Domain, error)
}
