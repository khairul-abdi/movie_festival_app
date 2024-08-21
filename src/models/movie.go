package models

import (
	"time"

	"github.com/lib/pq"
)

type (
	Movie struct {
		ID          uint           `gorm:"primaryKey" json:"id"`
		Title       string         `gorm:"type:varchar;not null;column:title" valid:"Required" json:"title"`
		Description string         `gorm:"type:text;column:description" json:"description"`
		Duration    string         `gorm:"type:varchar;not null;column:duration" valid:"Required" json:"duration"`
		Artists     pq.StringArray `gorm:"type:varchar[];not null;column:artists" valid:"Required" json:"artists"`
		Genres      pq.StringArray `gorm:"type:varchar[];not null;column:genres" valid:"Required" json:"genres"`
		WatchUrl    string         `gorm:"type:varchar;not null;column:watch_url" valid:"Required" json:"watch_url"`
		Viewed      pq.Int64Array  `gorm:"type:integer[];column:viewed" json:"viewed"`
		Voted       pq.Int64Array  `gorm:"type:integer[];column:voted" json:"voted"`
		CreatedAt   time.Time      `gorm:"column:created_at"  json:"created_at"`
		UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	}

	ParamMovie struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Artist      string `json:"artist"`
		Genre       string `json:"genre"`
	}
)
