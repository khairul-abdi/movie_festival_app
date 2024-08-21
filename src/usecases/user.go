package usecases

import (
	"movie_festival_app/packages"
	"movie_festival_app/src/helpers"
	"movie_festival_app/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *uc) UserRegister(c *gin.Context) (map[string]interface{}, string, int) {
	User := models.User{}
	c.ShouldBindJSON(&User)

	message, statusCode := helpers.ValidationForm(&User)
	if statusCode != 200 {
		return nil, message, statusCode
	}

	err := u.repo.InsertOne(&User)
	if err != nil {
		code, message := packages.Validate(err)
		return nil, message, code
	}

	data := map[string]interface{}{
		"id":       User.Id,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
		"role":     User.Role,
	}

	return data, "Success insert data", http.StatusCreated
}

func (u *uc) UserLogin(c *gin.Context) (map[string]interface{}, string, int) {
	User := models.User{}
	c.ShouldBindJSON(&User)

	password := User.Password
	whereVariable := "email = ?"
	whereValue := User.Email

	err := u.repo.FindOne(&User, whereVariable, whereValue)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, "email not registered", http.StatusUnauthorized
		}
		return nil, "invalid email", http.StatusUnauthorized
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		return nil, "Wrong password", http.StatusUnauthorized
	}

	data := map[string]interface{}{
		"token": helpers.GenerateToken(User.Id, User.Email, User.Role),
	}

	return data, "Success Create Token", http.StatusOK
}

func (u *uc) UpdateUser(c *gin.Context, userID uint) (map[string]interface{}, string, int) {

	var (
		newUser models.UserUpdateRequest
		oldUser models.User
		entity  models.User
	)
	c.ShouldBindJSON(&newUser)

	message, statusCode := helpers.ValidationForm(&newUser)
	if statusCode != 200 {
		return nil, message, statusCode
	}

	whereVariable := "id = ?"
	whereValue := userID
	err := u.repo.FindOne(&oldUser, whereVariable, whereValue)
	if err != nil {
		return nil, "Data Not Found", http.StatusBadRequest
	}

	entity.Id = userID
	entity.Username = newUser.Username
	entity.Email = newUser.Email
	entity.Password = oldUser.Password
	entity.Age = oldUser.Age
	entity.CreatedAt = oldUser.CreatedAt
	entity.UpdatedAt = time.Now()

	err = u.repo.UpdateUser(&oldUser, &entity)
	if err != nil {
		return nil, "Bad Request", http.StatusBadRequest
	}

	data := map[string]interface{}{
		"id":         oldUser.Id,
		"email":      oldUser.Email,
		"username":   oldUser.Username,
		"age":        oldUser.Age,
		"updated_at": oldUser.UpdatedAt,
	}

	return data, "Success Update user", http.StatusOK
}

func (u *uc) DeleteUser(c *gin.Context, userID int) error {
	User := models.User{}
	err := u.repo.DeleteUser(userID, &User)
	if err != nil {
		return err
	}

	return nil
}
