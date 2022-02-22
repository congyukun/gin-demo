package admin

import (
	"gin_demo/app/admin/Controller"
	"github.com/gin-gonic/gin"
)

var user = Controller.UserController{}

func Routers(e *gin.Engine) {
	r := e.Group("/admin")
	r.POST("/user/login", user.Login)
	r.GET("/user/index", user.Index)
	// e.GET("/goods", goodsHandler)
	// e.GET("/checkout", checkoutHandler)

}
