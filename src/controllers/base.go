package controllers

import (
	"movie_festival_app/src/usecases"

	"github.com/gin-gonic/gin"
)

type ctrl struct {
	uc usecases.UC
}

type Controllers interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	GetMovies(c *gin.Context)
	StoreMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	GetMoviesMostViewed(c *gin.Context)
	SearchMovies(c *gin.Context)
	TrackMovieViewership(c *gin.Context)
	VoteMovie(c *gin.Context)
	UnVoteMovie(c *gin.Context)
}

func NewCtrl(uc usecases.UC) Controllers {
	return &ctrl{uc: uc}
}
