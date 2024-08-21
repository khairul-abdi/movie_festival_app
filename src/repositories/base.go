package repositories

import (
	"movie_festival_app/src/models"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	InsertOne(data interface{}) error
	FindOne(data interface{}, whereVariable string, whereValue interface{}) error

	UpdateUser(oldUser, entity *models.User) error
	DeleteUser(id int, User *models.User) error

	FindAllMovies(models.ParamMovie) ([]models.Movie, error)
	UpdateMovie(movieId int, movie *models.Movie) error
	FindAllMoviesMostViewed() ([]models.Movie, error)
	SearchMovies(page, limit int) ([]models.Movie, error)
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
