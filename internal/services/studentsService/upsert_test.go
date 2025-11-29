package studentsService

import (
	"context"
	"errors"
	"testing"

	"github.com/Domenick1991/students/internal/models"
	"github.com/Domenick1991/students/internal/services/studentsService/mocks"
	"github.com/stretchr/testify/suite"
	"gotest.tools/v3/assert"
)

type StudentServiceSuite struct {
	suite.Suite
	ctx            context.Context
	studentStorage *mocks.StudentStorage
	studentService *StudentService
}

func (s *StudentServiceSuite) SetupTest() {
	s.studentStorage = mocks.NewStudentStorage(s.T())
	s.ctx = context.Background()
	s.studentService = NewStudentService(s.ctx, s.studentStorage, 0, 100)
}

func (s *StudentServiceSuite) TestUpsertSuccess() {
	studentsInfos := []*models.StudentInfo{
		{
			ID:    1,
			Name:  "Vasya",
			Age:   25,
			Email: "vasya@mail.ru",
		},
	}

	s.studentStorage.EXPECT().UpsertStudentInfo(s.ctx, studentsInfos).Return(nil)

	err := s.studentService.UpsertStudentInfo(s.ctx, studentsInfos)

	assert.NilError(s.T(), err)

}

func (s *StudentServiceSuite) TestUpsertError() {
	studentsInfos := []*models.StudentInfo{
		{
			ID:    1,
			Name:  "Vasya",
			Age:   25,
			Email: "vasya@mail.ru",
		},
	}
	wantErr := errors.New("error")

	s.studentStorage.EXPECT().UpsertStudentInfo(s.ctx, studentsInfos).Return(wantErr)

	err := s.studentService.UpsertStudentInfo(s.ctx, studentsInfos)

	assert.ErrorIs(s.T(), err, wantErr)

}

func (s *StudentServiceSuite) TestUpsertNameError() {
	studentsInfos := []*models.StudentInfo{
		{
			ID:    1,
			Name:  "",
			Age:   25,
			Email: "vasya@mail.ru",
		},
	}

	err := s.studentService.UpsertStudentInfo(s.ctx, studentsInfos)

	assert.Check(s.T(), err != nil)

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(StudentServiceSuite))
}
