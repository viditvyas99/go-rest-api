package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Go-REST-API/internal/storage"
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
			response.WriteJSON(w, http.StatusBadRequest,
				response.GenerateErrorResponse(err, http.StatusBadRequest))
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

		store := storage.NewMySQLStorage()

		id, err := store.CreateStudent(student)
		if err != nil {
			slog.Error("Failed to create student", "error", err)
			response.WriteJSON(w, http.StatusInternalServerError,
				response.GenerateErrorResponse(err, http.StatusInternalServerError))
			return
		}

		slog.Info("Student created successfully", "id", id)

		response.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Student created successfully", "id": string(id)})

	}
}

func GetStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to get student by ID")

		//get id from url path

		idStr := r.PathValue("id")

		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateErrorResponse(errors.New("invalid id format"), http.StatusBadRequest))
			return
		}

		store := storage.NewMySQLStorage()

		student, err := store.GetStudentByID(id)
		if err != nil {
			slog.Error("Failed to get student", "error", err)
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateErrorResponse(err, http.StatusInternalServerError))
			return
		}

		response.WriteJSON(w, http.StatusOK, student)
	}
}

func GetList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to get list of students")

		store := storage.NewMySQLStorage()

		students, err := store.GetAllStudents()
		if err != nil {
			slog.Error("Failed to get students", "error", err)
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateErrorResponse(err, http.StatusInternalServerError))
			return
		}

		response.WriteJSON(w, http.StatusOK, students)
	}
}

func UpdateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to update student")

		idStr := r.PathValue("id")

		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateErrorResponse(errors.New("invalid id format"), http.StatusBadRequest))
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest,
				response.GenerateErrorResponse(err, http.StatusBadRequest))
			return
		}

		validate := validator.New()
		err = validate.Struct(student)
		if err != nil {
			slog.Error("Validation failed", "error", err)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))

			return
		}

		store := storage.NewMySQLStorage()

		err = store.UpdateStudent(id, student)
		if err != nil {
			slog.Error("Failed to update student", "error", err)
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateErrorResponse(err, http.StatusInternalServerError))
			return
		}

		response.WriteJSON(w, http.StatusOK, map[string]string{"message": "Student updated successfully", "id": string(id)})
	}
}

func DeleteStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to delete student")

		idStr := r.PathValue("id")

		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateErrorResponse(errors.New("invalid id format"), http.StatusBadRequest))
			return
		}

		store := storage.NewMySQLStorage()

		err = store.DeleteStudent(id)
		if err != nil {
			slog.Error("Failed to delete student", "error", err)
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateErrorResponse(err, http.StatusInternalServerError))
			return
		}

		response.WriteJSON(w, http.StatusOK, map[string]string{"message": "Student deleted successfully"})
	}
}
