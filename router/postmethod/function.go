package postmethod

import (
	"fmt"
	"ginmod1/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
	Role     string
}

func (User) TableName() string {
	return "userinfo"
}

func Index(ctx *gin.Context) {
	username := ctx.PostForm("username")

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "首頁",
		"username": username,
	})
}

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	var userInfo = User{}
	model.UserDb.Where("username = ?", username).First(&userInfo)

	if userInfo.Password == password {
		session := sessions.Default(ctx)
		session.Options(sessions.Options{MaxAge: 300})
		session.Set("username", userInfo.Username)
		session.Set("role", userInfo.Role)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "首頁",
			"username": userInfo.Username,
			"role":     userInfo.Role,
		})
	} else {
		fmt.Println("失敗")
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title":     "首頁",
			"faillogin": true,
		})
	}

}

func DoAddUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	var userInfo = User{}
	userInfo.Username = username
	userInfo.Password = password
	userInfo.Role = "normal"
	model.UserDb.Create(&userInfo)
	ctx.String(http.StatusOK, "更新完成")
}

func modifyUserButton(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")
	ctx.HTML(http.StatusOK, "modifyUser.html", gin.H{
		"title":    "首頁",
		"username": username,
	})
}

func ModifyUserPassword(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")

	password := ctx.PostForm("password")
	model.UserDb.Model(&User{}).Where("username = ?", username).Update("password", password)
	ctx.String(http.StatusOK, "更新完成")
}
