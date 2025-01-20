package service

import "github.com/NCorreia100/GoAPI/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
