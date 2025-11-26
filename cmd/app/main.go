package main

import (
	"fmt"
	"os"

	"github.com/Domenick1991/students/config"
	"github.com/Domenick1991/students/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	studentsStorage := bootstrap.InitPGStorage(cfg)
	studentService := bootstrap.InitStudentService(studentsStorage)
	studentsInfoProcessor := bootstrap.InitStudentsInfoProcessor(studentService)
	studentsinfoupsertconsumer := bootstrap.InitStudentInfoUpsertConsumer(cfg, studentsInfoProcessor)
	students_api := bootstrap.InitStudentServiceAPI(studentService)

	bootstrap.AppRun(*students_api, studentsinfoupsertconsumer)
}
