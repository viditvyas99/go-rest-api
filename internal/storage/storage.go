package storage

import "github.com/Go-REST-API/internal/types"

type Storage interface {
	// Define methods for storage operations, e.g., CRUD operations for students
	CreateStudent(student types.Student) (int64, error)
}
