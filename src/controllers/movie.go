package controllers

import (
	"fmt"
	"movie_festival_app/packages"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *ctrl) GetMovies(c *gin.Context) {
	res, message, code := ctrl.uc.GetMovies(c)
	packages.Response(c, message, code, res)
}

func (ctrl *ctrl) StoreMovie(c *gin.Context) {
	fmt.Println("MASUKKKKKK")
	message, code := ctrl.uc.StoreMovie(c)
	packages.Response(c, message, code, nil)
}

func (ctrl *ctrl) UpdateMovie(c *gin.Context) {
	movieIdInt, err := strconv.Atoi(c.Param("movieId"))
	if err != nil {
		packages.Response(c, err.Error(), http.StatusInternalServerError, nil)
	}

	message, code := ctrl.uc.UpdateMovie(c, movieIdInt)
	packages.Response(c, message, code, nil)
}

func (ctrl *ctrl) GetMoviesMostViewed(c *gin.Context) {
	res, message, code := ctrl.uc.GetMoviesMostViewed(c)
	packages.Response(c, message, code, res)
}

func (ctrl *ctrl) SearchMovies(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		packages.Response(c, err.Error(), http.StatusInternalServerError, nil)
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		packages.Response(c, err.Error(), http.StatusInternalServerError, nil)
	}

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}
	fmt.Println("PAGE:", page)
	fmt.Println("LIMIT:", limit)

	res, message, code := ctrl.uc.SearchMovies(c, page, limit)
	packages.Response(c, message, code, res)
}
