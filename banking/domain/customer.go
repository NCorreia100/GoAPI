package domain

type Customer struct {
	Id      string
	Name    string
	City    string
	Zipcode string
	DOB     string
	Status  int16
}

type CustomerRepo interface {
	FindAll() ([]Customer, error)
}
