package bootstrap

import (
	"fmt"

	"github.com/Domenick1991/students/config"
	studentsinfoupsertconsumer "github.com/Domenick1991/students/internal/consumer/students_Info_upsert_consumer"
	studentsinfoprocessor "github.com/Domenick1991/students/internal/services/processors/students_info_processor"
)

func InitStudentInfoUpsertConsumer(cfg *config.Config, studentsInfoProcessor *studentsinfoprocessor.StudentsInfoProcessor) *studentsinfoupsertconsumer.StudentInfoUpsertConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return studentsinfoupsertconsumer.NewStudentInfoUpsertConsumer(studentsInfoProcessor, kafkaBrockers, cfg.Kafka.StudentInfoUpsertTopicName)
}
