/*
Create time at 2023年3月18日0018下午 20:35:28
Create User at Administrator
*/

package routes

import (
	"JWT-Test-Golang/controllers"
	"JWT-Test-Golang/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func Initrouter() {
	r := gin.Default()
	r.POST("/Signup", controllers.Signup)
	r.POST("/Login", controllers.Login)
	r.GET("/Validate", middleware.ParseToken, controllers.Validate)
	r.Run(os.Getenv("PORT"))
}
