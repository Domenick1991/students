package bootstrap

import (
	"context"

	"github.com/Domenick1991/students/internal/services/students"
	"github.com/Domenick1991/students/internal/storage/pgstorage"
)

func NewStudentService(storage *pgstorage.PGstorage) *students.StudentService {
	return students.NewStudentService(context.Background(), storage)
}
