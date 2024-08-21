package usecases

import (
	"movie_festival_app/packages"
	"movie_festival_app/src/helpers"
	"movie_festival_app/src/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (u *uc) GetMovies(c *gin.Context) (movies []models.Movie, statusMessage string, statusCode int) {
	var (
		param = models.ParamMovie{}
	)

	err := c.ShouldBindJSON(&param)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	res, err := u.repo.FindAllMovies(param)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return res, models.Success, http.StatusOK
}

func (u *uc) StoreMovie(c *gin.Context) (string, int) {
	movie := models.Movie{}
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	message, statusCode := helpers.ValidationForm(movie)
	if statusCode != 200 {
		return message, statusCode
	}

	err = u.repo.InsertOne(&movie)
	if err != nil {
		code, message := packages.Validate(err)
		return message, code
	}

	return models.Success, http.StatusOK
}

func (u *uc) UpdateMovie(c *gin.Context, movieId int) (string, int) {
	movie := models.Movie{}
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	message, statusCode := helpers.ValidationForm(movie)
	if statusCode != 200 {
		return message, statusCode
	}

	whereVariable := "id = ?"
	whereValue := movieId
	err = u.repo.FindOne(&models.Movie{}, whereVariable, whereValue)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	err = u.repo.UpdateMovie(movieId, &movie)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	return models.Success, http.StatusOK
}

func (u *uc) GetMoviesMostViewed(c *gin.Context) ([]models.Movie, string, int) {
	res, err := u.repo.FindAllMoviesMostViewed()
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return res, models.Success, http.StatusOK
}

func (u *uc) SearchMovies(c *gin.Context, page, limit int) ([]models.Movie, string, int) {
	res, err := u.repo.SearchMovies(page, limit)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return res, models.Success, http.StatusOK
}

func (u *uc) TrackMovieViewership(c *gin.Context) ([]models.Movie, string, int) {
	res, err := u.repo.TrackMovieViewership()
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return res, models.Success, http.StatusOK
}

func (u *uc) VoteMovie(c *gin.Context, movieId int) (string, int) {
	movie, statusMessage, statusCode := u.GetMovie(c, movieId)
	if statusCode != 200 {
		return statusMessage, statusCode
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	for i := 0; i < len(movie.Voted); i++ {
		if movie.Voted[i] != int64(userId) {
			continue
		} else {
			return "You already vote this movie", http.StatusOK
		}
	}

	movie.Voted = append(movie.Voted, int64(userId))

	err := u.repo.VoteMovie(movieId, movie)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	return models.Success, http.StatusOK
}

func (u *uc) GetMovie(c *gin.Context, movieId int) (movie *models.Movie, statusMessage string, statusCode int) {
	whereVariable := "id = ?"
	whereValue := movieId
	err := u.repo.FindOne(&movie, whereVariable, whereValue)
	if err != nil {
		return nil, err.Error(), http.StatusInternalServerError
	}

	return movie, models.Success, http.StatusOK
}

func (u *uc) UnVoteMovie(c *gin.Context, movieId int) (string, int) {
	movie, statusMessage, statusCode := u.GetMovie(c, movieId)
	if statusCode != 200 {
		return statusMessage, statusCode
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	isExistVote := false
	indexVoted := 0

	for i := 0; i < len(movie.Voted); i++ {
		if movie.Voted[i] == int64(userId) {
			isExistVote = true
			indexVoted = i
			break
		}
	}

	if !isExistVote {
		return "You do not vote this movie", http.StatusOK
	}

	movie.Voted = helpers.RemoveElement(indexVoted, movie.Voted)

	err := u.repo.VoteMovie(movieId, movie)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	return models.Success, http.StatusOK
}
