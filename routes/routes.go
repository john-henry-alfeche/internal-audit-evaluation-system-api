package routes

import (
	"github.com/gin-gonic/gin"
	"internal-audit-evaluation-system-api/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
	}
}
