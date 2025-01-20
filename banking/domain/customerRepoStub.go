package domain

type CustomerRepoStub struct {
	customers []Customer
}

func (s CustomerRepoStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepoStub() CustomerRepoStub {
	customers := []Customer{
		{Id: "1", Name: "Anh", City: "Hanoi", Zipcode: "100000", DOB: "1992-01-01", Status: 1},
		{Id: "2", Name: "Ram", City: "New Jersey", Zipcode: "08910", DOB: "1992-01-01", Status: 1},
	}
	return CustomerRepoStub{customers}
}
