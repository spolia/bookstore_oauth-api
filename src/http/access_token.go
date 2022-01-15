package http

import (
	"net/http"

	"github.com/bookstore_oauth-api/src/domain/access_token"
	access_token2 "github.com/bookstore_oauth-api/src/services/access_token"
	"github.com/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token2.Service
}

func NewHandler(service access_token2.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	at, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest("invalid json body"))
		return
	}

	if _, err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}
