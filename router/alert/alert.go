package alert

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlertRoute(r *gin.Engine) {
	getmethod := r.Group("/alert")
	{
		getmethod.POST("/", Index)

	}
}

func Index(ctx *gin.Context) {
	var message AlertMessage
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		ctx.String(http.StatusBadRequest, "傳送失敗", err)
		fmt.Println("errrr1")
		return
	}
	serviceLabel := GetServiceInfo(message)
	fmt.Println("errrr2")
	err = ScaleService(serviceLabel)
	fmt.Println("errrr3")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"request": "ok",
	})
}
