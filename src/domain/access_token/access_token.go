package access_token

import (
	"fmt"
	"github.com/bookstore_oauth-api/src/utils/crypto_utils"
	"strings"
	"time"

	"github.com/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json: "access_token"`
	UserID      int64  `json: "user_id"`
	ClientID    int64  `json: "client_id"`
	Expires     int64  `json: "expires"`
}

type AccessTokenRequest struct {
	GrantType string `json: "grant_type" validate:"required"`
	Scope     string `json: "scope"`

	// used for password grant type
	Username string `json: "username" validate:"required"`
	Password string `json: "password" validate:"required"`

	// used for client credentials grant type
	ClientID     string `json: "client_id"`
	ClientSecret string `json: "client_secret"`
}

func NewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserID:  userId,
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

func (at *AccessTokenRequest) Validate() *errors.RestError {
	if at.GrantType != grantTypePassword || at.GrantType != grantTypeClientCredentials {
		return errors.NewBadRequest("invalid grant type")
	}
	return nil
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
