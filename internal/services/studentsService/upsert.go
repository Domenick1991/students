package studentsService

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
)

func (s *StudentService) UpsertStudentInfo(ctx context.Context, studentsInfos []*models.StudentInfo) error {

	if err := s.validateInfo(studentsInfos); err != nil {
		return err
	}
	return s.studentStorage.UpsertStudentInfo(ctx, studentsInfos)
}
