package repositories

import (
	"movie_festival_app/src/models"
)

func (r *repo) DeleteUser(id int, User *models.User) error {
	err := r.db.Delete(&User, id).Error
	if err != nil {
		return err
	}

	return nil
}
