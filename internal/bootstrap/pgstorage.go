package bootstrap

import (
	"fmt"
	"log"

	"github.com/Domenick1991/students/config"
	"github.com/Domenick1991/students/internal/storage/pgstorage"
)

func InitPGStorage(cfg *config.Config) *pgstorage.PGstorage {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	storage, err := pgstorage.NewPGStorge(connectionString)
	if err != nil {
		log.Panic(fmt.Sprintf("ошибка инициализации БД, %v", err))
		panic(err)
	}
	return storage
}
