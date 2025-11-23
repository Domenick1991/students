package bootstrap

import (
	server "github.com/Domenick1991/students/internal/api/students_api"
	"github.com/Domenick1991/students/internal/services/students"
)

func NewStudentServiceAPI(studentService *students.StudentService) *server.StudentServiceAPI {
	return server.NewStudentServiceAPI(studentService)
}
