package repositories

import (
	"movie_festival_app/src/models"
)

func (r *repo) UpdateUser(oldUser, entity *models.User) error {
	err := r.db.Model(&oldUser).Updates(entity).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateMovie(movieId int, movie *models.Movie) error {
	err := r.db.Model(&movie).Where("id = ?", movieId).Updates(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) VoteMovie(movieId int, movie *models.Movie) error {
	err := r.db.Model(&models.Movie{}).Where("id = ?", movieId).Updates(movie).Error
	if err != nil {
		return err
	}
	return nil
}
