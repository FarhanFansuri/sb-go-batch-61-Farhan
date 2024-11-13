package routers

import (
	"formative-13/controllers"

	"github.com/gin-gonic/gin"
)

func RouterCar(app *gin.Engine) {
	router := app
	router.GET("/cars", controllers.GetCar)
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carId", controllers.UpdateCar)
	router.DELETE("/cars/:carId", controllers.DeleteCar)
}
