package service

import (
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

type TeamService struct {
	repo repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) CreateTeam(team model.Team) (int, error) {
	return s.repo.CreateTeam(team)
}

func (s *TeamService) GetTeamByID(id int) (model.Team, error) {
	return s.repo.GetTeamByID(id)
}

func (s *TeamService) GetTeamsByCompanyID(companyID int) ([]model.Team, error) {
	return s.repo.GetTeamsByCompanyID(companyID)
}

func (s *TeamService) DeleteTeam(id int) error {
	return s.repo.DeleteTeam(id)
}
