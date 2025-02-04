package utils

import (
	"github.com/gin-gonic/gin"
)

func BindValidate(c *gin.Context, dto interface{}) {
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(err)
	}
}
