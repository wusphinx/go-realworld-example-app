package apis

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	router.POST("/testing", startPage)
}
