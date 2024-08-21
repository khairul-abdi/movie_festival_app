package repositories

func (r *repo) InsertOne(data interface{}) error {
	err := r.db.Debug().Create(data).Error
	if err != nil {
		return err
	}

	return nil
}
