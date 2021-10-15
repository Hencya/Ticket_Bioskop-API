package addressesEntity

type Domain struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Street   string `son:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type Repository interface {
	Insert(address *Domain) (Domain, error)
	FindByID(id uint) (Domain, error)
	FindByCity(city string) ([]Domain, error)
}
