package app

import (
	"github.com/bookstore_oauth-api/src/clients/cassandra"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/http"
	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	if err := cassandra.GetSession(); err != nil {
		panic(err)
	}

	atHandler := http.NewHandler(access_token.NewService(db.New()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
