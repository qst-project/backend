package service

import (
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/repository"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
)


type Survey interface {
	GetSurvey(id string) (api.Survey, error)
	SetSurvey(servey api.Survey) (bool, string, error)
	// Update
	// Delete
}

type Service struct {
	Survey
}
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Survey: NewSurveyService(repos),
	}
}