package bootstrap

import (
	studentsinfoprocessor "github.com/Domenick1991/students/internal/services/processors/students_info_processor"
	"github.com/Domenick1991/students/internal/services/studentsService"
)

func InitStudentsInfoProcessor(studentService *studentsService.StudentService) *studentsinfoprocessor.StudentsInfoProcessor {
	return studentsinfoprocessor.NewStudentsInfoProcessor(studentService)
}
