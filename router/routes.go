package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	routes := r.Group("")
	{
		routes.GET("/genesisChallenge", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hi Genesis Bank",
			})
		})
	}

}
