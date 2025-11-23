package studentsinfoupsertconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/Domenick1991/students/internal/models"
	"github.com/segmentio/kafka-go"
)

type studentsInfoProcessor interface {
	Handle(ctx context.Context, studentsInfo *models.StudentInfo) error
}

type StudentInfoUpsertConsumer struct {
	studentsInfoProcessor studentsInfoProcessor
	kafkaBroker           []string
	topicName             string
}

func (c *StudentInfoUpsertConsumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           []string{"localhost:9092", "kafka:9094"},
		GroupID:           "StudentService_group",
		Topic:             c.topicName,
		//StartOffset:       kafka.FirstOffset,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("StudentInfoUpsertConsumer.consume error", "error", err.Error())
		}
		var studentInfo *models.StudentInfo
		err = json.Unmarshal(msg.Value, &studentInfo)
		if err != nil {
			slog.Error("parce", "error", err)
			continue
		}
		err = c.studentsInfoProcessor.Handle(ctx, studentInfo)
		if err != nil {
			slog.Error("Handle", "error", err)
		}
	}

}

func NewStudentInfoUpsertConsumer(studentsInfoProcessor studentsInfoProcessor, kafkaBroker []string, topicName string) *StudentInfoUpsertConsumer {
	return &StudentInfoUpsertConsumer{
		studentsInfoProcessor: studentsInfoProcessor,
		kafkaBroker:           kafkaBroker,
		topicName:             topicName,
	}
}
