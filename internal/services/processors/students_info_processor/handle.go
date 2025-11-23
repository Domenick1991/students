package studentsinfoprocessor

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
)

func (p *StudentsInfoProcessor) Handle(ctx context.Context, studentsInfo *models.StudentInfo) error {
	return p.studentService.UpsertStudentInfo(ctx, []*models.StudentInfo{studentsInfo})
}
