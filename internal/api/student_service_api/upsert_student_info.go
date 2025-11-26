package student_service_api

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
	proto_models "github.com/Domenick1991/students/internal/pb/models"
	"github.com/Domenick1991/students/internal/pb/students_api"
	"github.com/samber/lo"
)

func (s *StudentServiceAPI) UpsertStudentInfos(ctx context.Context, req *students_api.GetStudentInfoUpsertRequest) (*students_api.GetStudentInfoUpsertResponce, error) {
	err := s.studentService.UpsertStudentInfo(ctx, mapStudentInfo(req.Students))
	if err != nil {
		return &students_api.GetStudentInfoUpsertResponce{}, err
	}
	return &students_api.GetStudentInfoUpsertResponce{}, nil
}

func mapStudentInfo(studentInfo []*proto_models.StudentsUpsertModel) []*models.StudentInfo {
	return lo.Map(studentInfo, func(s *proto_models.StudentsUpsertModel, _ int) *models.StudentInfo {
		return &models.StudentInfo{
			Name:  s.Name,
			Age:   s.Age,
			Email: s.Email,
		}
	})
}
