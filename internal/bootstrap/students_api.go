package bootstrap

import (
	server "github.com/Domenick1991/students/internal/api/student_service_api"
	"github.com/Domenick1991/students/internal/services/studentsService"
)

func InitStudentServiceAPI(studentService *studentsService.StudentService) *server.StudentServiceAPI {
	return server.NewStudentServiceAPI(studentService)
}
