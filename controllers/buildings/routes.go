package buildings

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.Engine) {
	buildings := route.Group("/buildings")
	{
		buildings.GET("", Ping)
	}
}
