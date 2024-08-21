package models

import (
	"movie_festival_app/src/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id              uint      `gorm:"primaryKey" json:"id"`
	Username        string    `gorm:"type:varchar(20);not null;column:username;uniqueIndex" valid:"Required"`
	Email           string    `gorm:"type:varchar(100);not null;column:email;uniqueIndex" json:"email" valid:"Required; Email"`
	Password        string    `gorm:"type:varchar(100);not null;column:password" json:"password" valid:"Required; MinSize(6)"`
	Age             int       `gorm:"column:age;not null" valid:"Required; Min(9)"`
	Role            string    `gorm:"column:role"`
	TotalWatchMovie time.Time `gorm:"column:total_wath_movie"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

type UserUpdateRequest struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username" valid:"Required"`
	Email    string `json:"email" valid:"Required; Email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)
	return
}

func (r User) ToUserUpdateRequest() UserUpdateRequest {
	return UserUpdateRequest{
		Id:       int(r.Id),
		Username: r.Username,
		Email:    r.Email,
	}
}
