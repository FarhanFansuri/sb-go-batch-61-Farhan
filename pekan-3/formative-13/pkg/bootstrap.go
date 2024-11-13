package pkg

import (
	"formative-13/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Bootstrap() {
	godotenv.Load()

	port := os.Getenv("PORT")
	app := gin.Default()

	routers.RouterCar(app)
	app.Run(port)
}
