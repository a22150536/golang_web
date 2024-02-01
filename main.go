package main

import (
	"ginmod1/model"
	alertM "ginmod1/router/alert"

	// getM "ginmod1/router/getmethod"
	// postM "ginmod1/router/postmethod"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.Use(sessions.Sessions("gogin", model.UserSession))
	// getM.GetRoute(r)
	// postM.PostRoute(r)
	alertM.AlertRoute(r)
	r.Run("0.0.0.0:7777")
}
