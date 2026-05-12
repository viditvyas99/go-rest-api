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

func (m *MySQLStorage) GetAllStudents() ([]types.Student, error) {
	query := "SELECT id, name, email, class FROM students"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all students: %w", err)
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Class)
		if err != nil {
			return nil, fmt.Errorf("failed to scan student: %w", err)
		}
		students = append(students, student)
	}

	return students, nil
}

func (m *MySQLStorage) UpdateStudent(id int64, student types.Student) error {
	query := "UPDATE students SET name = ?, email = ?, class = ? WHERE id = ?"
	_, err := database.DB.Exec(query, student.Name, student.Email, student.Class, id)
	if err != nil {
		return fmt.Errorf("failed to update student: %w", err)
	}

	return nil
}

func (m *MySQLStorage) DeleteStudent(id int64) error {
	query := "DELETE FROM students WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}

	return nil
}
