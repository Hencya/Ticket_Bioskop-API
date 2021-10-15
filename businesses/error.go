package businesses

import "errors"

var (
	//General errors
	ErrInternalServer = errors.New("Something Gone Wrong,Please Contact Administrator")
	ErrNotFound       = errors.New("Data Not Found")
	ErrIdNotFound     = errors.New("Id Not Found")
	ErrDuplicateData  = errors.New("Data already exist")

	//Users errors
	ErrDuplicateEmail        = errors.New("Email already used")
	ErrEmailPasswordNotFound = errors.New("(Email) or (Password) empty")
	ErrEmailNotRegistered    = errors.New("Email not registered")
	ErrPassword              = errors.New("Wrong Password")
	ErrNotFoundUser          = errors.New("user doesn't exist")

	//Cinema errors
	ErrNearestCinemaNotFound = errors.New("No one Cinema found within your area")
	ErrCinemaNotFound        = errors.New("Cinema not found")
	ErrCinemaNotAvailable    = errors.New("Cinema is closed because of corona")
)
