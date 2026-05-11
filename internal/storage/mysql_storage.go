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
