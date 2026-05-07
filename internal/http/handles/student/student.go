package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Go-REST-API/internal/types"
	"github.com/Go-REST-API/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to create student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			slog.Error("Empty body", "error", err)

			response.WriteJSON(w, http.StatusBadRequest, response.GenerateErrorResponse(err, http.StatusBadRequest))

			return
		}
		//request body  valid json
		validate := validator.New()
		err = validate.Struct(student)
		if err != nil {
			slog.Error("Validation failed", "error", err)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))

			return
		}

		slog.Info("creating the student ")

		response.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})

	}
}
