package router

import (
	"github.com/danvieira97/genesisChallenge/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	basePath := "exchange"

	routes := r.Group(basePath)
	{
		routes.GET("/:amount/:from/:to/:rate", handler.CurrencyConverter)
		routes.GET("allCurrencyConverter", handler.AllCurrencyConverter)
	}

}
