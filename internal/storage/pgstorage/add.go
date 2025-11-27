package pgstorage

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (storage *PGstorage) UpsertStudentInfo(ctx context.Context, studentInfos []*models.StudentInfo) error {
	query := storage.upsertQuery(studentInfos)
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "generate query error")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "exeс query")
	}
	return err
}

func (storage *PGstorage) upsertQuery(studentInfos []*models.StudentInfo) squirrel.Sqlizer {
	infos := lo.Map(studentInfos, func(info *models.StudentInfo, _ int) *StudentInfo {
		return &StudentInfo{
			Name:  info.Name,
			Email: info.Email,
			Age:   info.Age,
		}
	})

	q := squirrel.Insert(tableName).Columns(NameСolumnName, EmailСolumnName, AgeСolumnName).
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		q = q.Values(info.Name, info.Email, info.Age)
	}
	return q
}
