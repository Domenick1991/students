package studentsinfoprocessor

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
)

type studentService interface {
	UpsertStudentInfo(ctx context.Context, studentsInfos []*models.StudentInfo) error
}

type StudentsInfoProcessor struct {
	studentService studentService
}

func NewStudentsInfoProcessor(studentService studentService) *StudentsInfoProcessor {
	return &StudentsInfoProcessor{
		studentService: studentService,
	}
}
