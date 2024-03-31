package utils

import (
	"encoding/json"
	"net/http"

	"github.com/sant470/grep-api/utils/errors"
)

func Decode(r *http.Request, v interface{}) *errors.AppError {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return errors.BadRequest(err.Error())
	}
	return nil
}
