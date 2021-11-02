package access_token

import (
	"strings"

	"github.com/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestError) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequest("invalid access token")
	}

	at, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return at, nil
}
