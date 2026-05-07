package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

//interface {} == any type

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	//encode data to json and write to response
	return json.NewEncoder(w).Encode(data)

}

func GenerateErrorResponse(err error, status int) ErrorResponse {
	return ErrorResponse{
		Error:  err.Error(),
		Status: status,
	}
}

func ValidationError(errs validator.ValidationErrors) ErrorResponse {

	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, err.Field()+" is required")
		case "email":
			errMsg = append(errMsg, err.Field()+" must be a valid email")
		default:
			errMsg = append(errMsg, err.Field()+" is invalid")
		}
	}

	return GenerateErrorResponse(errors.New(strings.Join(errMsg, ", ")), http.StatusBadRequest)
}
