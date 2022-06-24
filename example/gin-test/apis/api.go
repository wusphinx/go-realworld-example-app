package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "Success")
}
