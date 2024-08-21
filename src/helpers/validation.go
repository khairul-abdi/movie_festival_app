package helpers

import (
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
)

func ValidationForm(data interface{}) (message string, statusCode int) {
	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(data)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	if !b {
		// validation does not pass
		var message string
		for i, err := range valid.Errors {
			if i == 0 {
				message = err.Message
			} else {
				message += ", " + err.Message
			}
		}
		return message, http.StatusBadRequest
	}

	return "", http.StatusOK
}
