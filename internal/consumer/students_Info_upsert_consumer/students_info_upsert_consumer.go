package studentsinfoupsertconsumer

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
)

type studentsInfoProcessor interface {
	Handle(ctx context.Context, studentsInfo *models.StudentInfo) error
}

type StudentInfoUpsertConsumer struct {
	studentsInfoProcessor studentsInfoProcessor
	kafkaBroker           []string
	topicName             string
}

func NewStudentInfoUpsertConsumer(studentsInfoProcessor studentsInfoProcessor, kafkaBroker []string, topicName string) *StudentInfoUpsertConsumer {
	return &StudentInfoUpsertConsumer{
		studentsInfoProcessor: studentsInfoProcessor,
		kafkaBroker:           kafkaBroker,
		topicName:             topicName,
	}
}
