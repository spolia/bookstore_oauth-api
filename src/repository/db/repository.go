package db

import (
	"github.com/bookstore_oauth-api/src/clients/cassandra"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryAccessToken          = "SELECT access_token,user_id,client_id,expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken    = "INSERT INTO access_tokens(access_token,user_id,client_id,expires) VALUES(?,?,?,?);"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires = ? WHERE access_token=?;"
)

type Repository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func New() Repository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {
	var result access_token.AccessToken
	err := cassandra.GetSession().Query(queryAccessToken, id).Scan(
		&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFound("no access token found with given id")
		}

		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestError {
	err := cassandra.GetSession().Query(queryCreateAccessToken, at.Expires, at.AccessToken).Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	err := cassandra.GetSession().Query(queryUpdateExpirationTime, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
