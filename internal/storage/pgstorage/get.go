package pgstorage

import (
	"context"

	"github.com/Domenick1991/students/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (storage *PGstorage) GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.StudentInfo, error) {
	query := storage.getQuery(IDs)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate query error")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "quering error")
	}
	var students []*models.StudentInfo
	for rows.Next() {
		var s models.StudentInfo
		if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Age); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		students = append(students, &s)
	}
	return students, nil
}

func (storage *PGstorage) getQuery(IDs []uint64) squirrel.Sqlizer {
	q := squirrel.Select(IDСolumnName, NameСolumnName, EmailСolumnName, AgeСolumnName).From(tableName).
		Where(squirrel.Eq{IDСolumnName: IDs}).PlaceholderFormat(squirrel.Dollar)
	return q
}
