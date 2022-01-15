package access_token

import (
	"strings"

	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/bookstore_oauth-api/src/repository/rest"
	"github.com/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type service struct {
	restUsersRepo rest.Repository
	dbRepo        db.Repository
}

func NewService(usersRepo rest.Repository, dbRepo db.Repository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetByID(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequest("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetByID(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.Username)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.NewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
