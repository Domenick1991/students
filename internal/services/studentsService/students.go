package studentsService

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
)

type studentStorage interface {
	GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.StudentInfo, error)
	UpsertStudentInfo(ctx context.Context, studentInfos []*models.StudentInfo) error
}

type StudentService struct {
	studentStorage studentStorage
}

func NewStudentService(ctx context.Context, studentStorage studentStorage) *StudentService {
	return &StudentService{
		studentStorage: studentStorage,
	}
}
