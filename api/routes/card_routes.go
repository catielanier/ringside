package routes

import (
	"github.com/catielanier/ringside/controllers"
	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.Engine) {
	r.GET("/cards", controllers.GetAllCards())
}
