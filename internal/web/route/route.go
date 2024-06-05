package route

import (
	"backend/internal/web/controller"
	"backend/internal/web/middleware"
	"backend/pkg/env"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	mode := env.Get("GIN_MODE", "release")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	routers := gin.Default()

	routersV1 := routers.Group("/api/v1")

	routersV1.Use(middleware.Cors())

	routersV1.GET("/server_check", controller.ServerCheck)

	routersV1.GET("/users", controller.GetUsers)

	return routers
}
