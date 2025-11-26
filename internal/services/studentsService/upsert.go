package studentsService

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/Domenick1991/students/internal/models"
)

func (s *StudentService) UpsertStudentInfo(ctx context.Context, studentsInfos []*models.StudentInfo) error {

	if err := validateInfo(studentsInfos); err != nil {
		return err
	}
	return s.studentStorage.UpsertStudentInfo(ctx, studentsInfos)
}

func validateInfo(studentsInfos []*models.StudentInfo) error {
	for _, info := range studentsInfos {
		if len(info.Name) == 0 || len(info.Name) > 100{
			return errors.New("имя не должно быть пустым и не должно превышать 100 символов")
		}
		if info.Age <= 0 || info.Age > 100 {
			return fmt.Errorf("некорректный возвраст у студента %v", info.Age)
		}
		if !isValidEmail(info.Email) {
			return fmt.Errorf("некорректный email у студента %v", info.Age)
		}
	}
	return nil
}

func isValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if len(parts[1]) == 0 || len(parts[1]) > 253 {
		return false
	}

	return true
}
