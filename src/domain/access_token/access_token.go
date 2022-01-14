package access_token

import (
	"strings"
	"time"

	"github.com/bookstore_oauth-api/src/utils/errors"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json: "access_token"`
	UserID      int64  `json: "user_id"`
	ClientID    int64  `json: "client_id"`
	Expires     int64  `json: "expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequest("invalid access token")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequest("invalid user id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequest("invalid expiration time")
	}
	return nil
}
