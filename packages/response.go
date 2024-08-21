package packages

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, message string, code int, data interface{}) {
	if data == nil {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
		return
	}

	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
