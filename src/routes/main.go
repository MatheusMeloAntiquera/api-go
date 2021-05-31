package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func Run() {
	addUserRoutes()
	Router.Run("localhost:" + os.Getenv("APP_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
