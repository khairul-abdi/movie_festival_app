package middlewares

import (
	"fmt"
	"movie_festival_app/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		fmt.Println("Error: ", err)
		fmt.Println("verifyToken: ", verifyToken)
		if err != nil {
			message := "Unauthenticated"
			code := http.StatusUnauthorized
			c.AbortWithStatusJSON(code,
				gin.H{
					"code":    code,
					"message": message,
				})
			return
		}
		c.Set("userData", verifyToken)

		c.Next()
	}
}
