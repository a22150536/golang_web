package getmethod

import (
	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {

	getmethod := r.Group("/")
	{
		getmethod.GET("/", Index)
		getmethod.GET("account", Account)
		getmethod.GET("login", Login)
		getmethod.GET("addUserButton", AddUserButton)
		getmethod.GET("news", News)
		getmethod.GET("admin", Admin)

	}

}
