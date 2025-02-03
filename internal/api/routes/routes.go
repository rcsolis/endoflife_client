package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rcsolis/endoflife_client/internal/api/handlers"
)

func GetRoutes(router *gin.Engine) {
	group := router.Group("/api")

	group.GET("/", handlers.GetAll)
	group.GET("/versions/:name", handlers.GetVersions)
	group.GET("/:name/version/:version", handlers.GetDetails)

}
