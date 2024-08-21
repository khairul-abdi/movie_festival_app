package controllers

import (
	"movie_festival_app/packages"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (ctrl *ctrl) UserRegister(c *gin.Context) {
	data, message, code := ctrl.uc.UserRegister(c)
	packages.Response(c, message, code, data)
}

func (ctrl *ctrl) UserLogin(c *gin.Context) {
	data, message, code := ctrl.uc.UserLogin(c)
	packages.Response(c, message, code, data)
}

func (ctrl *ctrl) UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIDJWT := uint(userData["id"].(float64))

	if userIDJWT != uint(userId) {
		message := "User not match, only the user can update account"
		code := http.StatusOK
		packages.Response(c, message, code, nil)
		return
	}

	data, message, code := ctrl.uc.UpdateUser(c, uint(userId))

	packages.Response(c, message, code, data)
}

func (ctrl *ctrl) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userIDJWT := uint(userData["id"].(float64))

	if userIDJWT != uint(userId) {
		message := "User not match, only the user can delete account"
		code := http.StatusOK
		packages.Response(c, message, code, nil)
		return
	}

	err = ctrl.uc.DeleteUser(c, userId)
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
		return
	}

	message := "Your user has been successfully deleted"
	code := http.StatusOK

	packages.Response(c, message, code, nil)

}
