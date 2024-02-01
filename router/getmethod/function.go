package getmethod

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")
	if username == nil {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首頁",
		})
	} else {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "首頁",
			"username": username,
		})
	}

}

func Admin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")
	role := session.Get("role")

	ctx.HTML(http.StatusOK, "admin.html", gin.H{
		"title":    "首頁",
		"username": username,
		"role": role,
	})
}

func Account(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	ctx.HTML(http.StatusOK, "account.html", gin.H{
		"title":    "帳號登入",
		"username": username,
		"password": password,
	})
}

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func AddUserButton(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addUser.html", gin.H{})
}

func News(ctx *gin.Context) {
	msg := ctx.Query("type")
	new := Article{
		Title:   "標題",
		Content: "內容",
	}
	ctx.HTML(http.StatusOK, "news.html", gin.H{
		"news": new,
		"msg":  msg,
	})
}
