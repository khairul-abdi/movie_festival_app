package repositories

import (
	"fmt"
	"movie_festival_app/src/helpers"
	"movie_festival_app/src/models"

	"gorm.io/gorm/clause"
)

func (r *repo) FindOne(data interface{}, whereVariable string, whereValue interface{}) error {
	err := r.db.Debug().Where(whereVariable, whereValue).Take(data).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) FindAllMovies(param models.ParamMovie) (movies []models.Movie, err error) {
	var (
		db            = r.db
		rawWhereArray []string
		rawQuery      = "WHERE "
	)

	query := `
		SELECT 
			*
		FROM (
				SELECT 
					*, 
					unnest(artists) artist, 
					UNNEST (genres) genre
				FROM 
					public.movies
			) 
	`

	if len(param.Title) > 0 {
		rawWhereArray = append(rawWhereArray, fmt.Sprintf("LOWER(title) LIKE LOWER('%%%s%%')", param.Title))
	}

	if len(param.Description) > 0 {
		rawWhereArray = append(rawWhereArray, fmt.Sprintf("LOWER(description) LIKE LOWER('%%%s%%')", param.Description))
	}

	if len(param.Artist) > 0 {
		rawWhereArray = append(rawWhereArray, fmt.Sprintf("LOWER(artist) LIKE LOWER('%%%s%%')", param.Artist))
	}

	if len(param.Genre) > 0 {
		rawWhereArray = append(rawWhereArray, fmt.Sprintf("LOWER(genre) LIKE LOWER('%%%s%%')", param.Genre))
	}

	for i := 0; i < len(rawWhereArray); i++ {
		if i == 0 {
			rawQuery += rawWhereArray[i]
		} else {
			rawQuery += " AND " + rawWhereArray[i]
		}
	}

	db.Raw(query + rawQuery).Scan(&movies)
	return
}

func (r *repo) FindAllMoviesMostViewed() (movies []models.Movie, err error) {
	err = r.db.Preload(clause.Associations).Order("array_length(viewed, 1) DESC").Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return
}

func (r *repo) SearchMovies(page, limit int) (movies []models.Movie, err error) {
	err = r.db.Preload(clause.Associations).Scopes(helpers.NewPaginate(limit, page).PaginatedResult).Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return
}

func (r *repo) TrackMovieViewership() (movies []models.Movie, err error) {
	err = r.db.Preload(clause.Associations).Raw(`
		SELECT 
			m.*,
			array_length(m.viewed,1)
		FROM 
			public."movies" m
		ORDER BY array_length(m.viewed,1) DESC;
	`).Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return
}
