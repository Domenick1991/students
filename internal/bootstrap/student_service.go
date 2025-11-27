package bootstrap

import (
	"context"

	"github.com/Domenick1991/students/config"
	"github.com/Domenick1991/students/internal/services/studentsService"
	"github.com/Domenick1991/students/internal/storage/pgstorage"
)

func InitStudentService(storage *pgstorage.PGstorage, cfg *config.Config) *studentsService.StudentService {

	return studentsService.NewStudentService(context.Background(), storage, cfg.StudentServiceSettings.MinNameLen, cfg.StudentServiceSettings.MaxNameLen)
}
