package service

import (
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

type CompanyService struct {
	repo repository.Company
}

func NewCompanyService(repo repository.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(company model.Company) (int, error) {
	return s.repo.CreateCompany(company)
}

func (s *CompanyService) GetCompanyByID(id int) (model.Company, error) {
	return s.repo.GetCompanyByID(id)
}

func (s *CompanyService) GetCompaniesByUserID(userID int) ([]model.Company, error) {
	return s.repo.GetCompaniesByUserID(userID)
}

func (s *CompanyService) DeleteCompany(id int) error {
	return s.repo.DeleteCompany(id)
}
