package studentsService

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/Domenick1991/students/internal/models"
)

func (s *StudentService) validateInfo(studentsInfos []*models.StudentInfo) error {
	for _, info := range studentsInfos {
		if len(info.Name) <= s.minNameLen || len(info.Name) >= s.maxNameLen {
			return errors.New("имя не должно быть пустым и не должно превышать 100 символов")
		}
		if info.Age <= 0 || info.Age > 100 {
			return fmt.Errorf("некорректный возвраст у студента %v", info.Age)
		}
		if !s.isValidEmail(info.Email) {
			return fmt.Errorf("некорректный email у студента %v", info.Age)
		}
	}
	return nil
}

func (s *StudentService) isValidEmail(email string) bool {
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
