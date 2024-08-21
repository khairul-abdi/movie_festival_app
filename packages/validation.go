package packages

import (
	"net/http"
	"strings"
)

func Validate(err error) (int, string) {
	errSplit := strings.Split(err.Error(), " ")
	if errSplit[7] == `"idx_users_username"` {
		return http.StatusBadRequest, "Username Already Used"
	} else if errSplit[7] == `"idx_users_email"` {
		return http.StatusBadRequest, "Email Already Used"
	} else {
		return http.StatusInternalServerError, "There is a problem with the resource"
	}
}
