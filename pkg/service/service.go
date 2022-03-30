package service

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/repository"
)

type Survey interface {
	GetSurvey(ref string) (*api.Survey, error)
	SetSurvey(survey *api.Survey) (bool, string, error)
	// Update
	DeleteSurvey(ref string) (bool, error)
}

type Service struct {
	Survey
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Survey: NewSurveyService(repos),
	}
}