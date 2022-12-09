package service

import (
	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/pkg/repository"
	"github.com/gofrs/uuid"
)

type CustomerService struct {
	repo repository.Customer
}

func NewCustomerService(repo repository.Customer) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Create(firstName, lastName, birthday, referrerCode, phone, email string) (string, error) {
	return s.repo.Create(firstName, lastName, birthday, referrerCode, phone, email)
}

func (s *CustomerService) GetAll(page, limit int) ([]gopitman.Customer, error) {
	return s.repo.GetAll(page, limit)
}

func (s *CustomerService) GetById(id uuid.UUID) (gopitman.Customer, error) {
	return s.repo.GetById(id)
}

func (s *CustomerService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *CustomerService) Update(id uuid.UUID, input gopitman.UpdateCustomer) error {
	return s.repo.Update(id, input)
}

// func (s *CustomerService) SearchCustomer(data map[string]string) (uuid.UUID, bool) {
// 	return s.repo.SearchCustomer(data)
// }
