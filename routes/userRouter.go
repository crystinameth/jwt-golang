package routes

import(
	controller "github.com/crystinameth/jwt-golang/controllers"
	"github.com/crystinameth/jwt-golang/middleware"
	"github.com/gin-gonic.gin"

)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())  // cuz by now user will have a token , authentication needed
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
