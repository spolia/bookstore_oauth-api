package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"

	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserWhenTheAPIReturnTimeout(t *testing.T) {

}

func TestLoginUserWhenInvalidError(t *testing.T) {

}

func TestLoginUserWhenInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserWhenNoError(t *testing.T) {

}

