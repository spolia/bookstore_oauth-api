package db

import (
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func New() Repository {
	return &dbRepository{}
}
func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {

}
