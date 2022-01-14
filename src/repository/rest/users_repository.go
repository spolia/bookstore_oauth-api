package rest

import (
	"encoding/json"
	"time"

	"github.com/bookstore_oauth-api/src/domain/users"
	"github.com/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

type Repository interface {
	LoginUser(string) (*users.User, *errors.RestError)
}

type userRepository struct{}

func New() Repository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string) (*users.User, *errors.RestError) {
	restClient := rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}

	result := restClient.Post("/users/login", users.UserLoginRequest{Email: email})
	if result == nil || result.Response == nil {
		return nil, errors.NewInternalServerError("invalid response when trying to login user")
	}

	if result.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(result.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid response when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(result.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("invalid response when trying to login user ")
	}

	return &user, nil
}
