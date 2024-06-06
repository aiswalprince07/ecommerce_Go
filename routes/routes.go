package routes

import(
	"github.com/aiswalprince07/ecommerce_Go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup")
	incomingRoutes.POST("/users/login")
	incomingRoutes.POST("/admin/addproduct")
	incomingRoutes.GET("/users/productview")
	incomingRoutes.GET("/users/search")
}