package student_service_api

import (
	"context"
	"log"

	"github.com/Domenick1991/students/internal/models"
	proto_models "github.com/Domenick1991/students/internal/pb/models"
	"github.com/Domenick1991/students/internal/pb/students_api"
	"github.com/samber/lo"
)

func (s *StudentServiceAPI) GetStudentInfoByIDs(ctx context.Context, req *students_api.GetStudentInfoByIDsRequest) (*students_api.GetStudentInfoByIDsResponse, error) {
	log.Printf("Received request with IDs: %v", req.Ids)

	responce, err := s.studentService.GetStudentInfoByIDs(ctx, req.Ids)
	if err != nil {
		return &students_api.GetStudentInfoByIDsResponse{}, err
	}
	return &students_api.GetStudentInfoByIDsResponse{StudentInfos: mapStudentInfoByResponce(responce)}, nil
}

func mapStudentInfoByResponce(studentInfo []*models.StudentInfo) []*proto_models.StudentsInfoModel {
	return lo.Map(studentInfo, func(s *models.StudentInfo, _ int) *proto_models.StudentsInfoModel {
		return &proto_models.StudentsInfoModel{
			Id:    s.ID,
			Name:  s.Name,
			Email: s.Email,
			Age:   s.Age,
		}
	})
}
