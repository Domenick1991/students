package studentsService

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
)

func (s *StudentService) GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.StudentInfo, error) {
	return s.studentStorage.GetStudentInfoByIDs(ctx, IDs)
}
