package postmethod

import (
	"github.com/gin-gonic/gin"
)

func PostRoute(r *gin.Engine) {

	postMethod := r.Group("/")
	{
		postMethod.POST("/", Index)
		postMethod.POST("/doAddUser", DoAddUser)
		postMethod.POST("/modifyUserButton", modifyUserButton)
		postMethod.POST("/modifyUser", ModifyUserPassword)
		postMethod.POST("/login", Login)

	}

}
