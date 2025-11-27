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
	studentService := bootstrap.InitStudentService(studentsStorage, cfg)
	studentsInfoProcessor := bootstrap.InitStudentsInfoProcessor(studentService)
	studentsinfoupsertconsumer := bootstrap.InitStudentInfoUpsertConsumer(cfg, studentsInfoProcessor)
	studentsApi := bootstrap.InitStudentServiceAPI(studentService)

	bootstrap.AppRun(*studentsApi, studentsinfoupsertconsumer)
}
