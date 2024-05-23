package router

import (
	"web/context"
	"web/web/handler"
	"web/web/logic/ping"

	"github.com/gin-gonic/gin"
)

// validator web server router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// set request start
	r.Use(func(ctx *gin.Context) {
		context.SetRequestTIme(ctx)
	})

	api := r.Group("/api")

	// server test
	pingGroup := api.Group("ping")
	{
		pingGroup.GET("", handler.TRPathParamHandler(ping.GetPingInfo))
	}

	return r
}
