package storage

import (
	"fmt"

	"github.com/Go-REST-API/database"
	"github.com/Go-REST-API/internal/types"
)

type MySQLStorage struct{}

func NewMySQLStorage() *MySQLStorage {
	return &MySQLStorage{}
}

func (m *MySQLStorage) CreateStudent(student types.Student) (int64, error) {

	query := "INSERT INTO students (name, email, class) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, student.Name, student.Email, student.Class)
	if err != nil {
		return 0, fmt.Errorf("failed to create student: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (m *MySQLStorage) GetStudentByID(id int64) (*types.Student, error) {
	query := "SELECT id, name, email, class FROM students WHERE id = ?"
	row := database.DB.QueryRow(query, id)

	var student types.Student
	err := row.Scan(&student.ID, &student.Name, &student.Email, &student.Class)
	if err != nil {
		return nil, fmt.Errorf("failed to get student by ID: %w", err)
	}

	return &student, nil
}
