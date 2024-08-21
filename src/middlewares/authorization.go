package middlewares

import (
	"movie_festival_app/src/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Admin role
		userData := c.MustGet("userData").(jwt.MapClaims)
		role := userData["role"]

		// compare
		if role != models.RoleAdmin {
			message := "Access Forbidden"
			code := http.StatusForbidden
			c.AbortWithStatusJSON(code,
				gin.H{
					"code":    code,
					"message": message,
				})
			return
		}

		c.Next()
	}
}
