package usecases

import (
	"movie_festival_app/src/models"
	"movie_festival_app/src/repositories"

	"github.com/gin-gonic/gin"
)

type uc struct {
	repo repositories.Repo
}

type UC interface {
	UserRegister(c *gin.Context) (map[string]interface{}, string, int)
	UserLogin(c *gin.Context) (map[string]interface{}, string, int)
	UpdateUser(c *gin.Context, userId uint) (map[string]interface{}, string, int)
	DeleteUser(c *gin.Context, userId int) error

	GetMovies(c *gin.Context) ([]models.Movie, string, int)
	StoreMovie(c *gin.Context) (string, int)
	UpdateMovie(c *gin.Context, movieId int) (string, int)
	GetMoviesMostViewed(c *gin.Context) ([]models.Movie, string, int)
	SearchMovies(c *gin.Context, page, limit int) ([]models.Movie, string, int)
}

func NewUC(repo repositories.Repo) UC {
	return &uc{repo: repo}
}
