package pgstorage

type StudentInfo struct {
	ID    uint64 `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Age   uint64 `db:"age"`
}

const (
	tableName = "studentsInfo"

	IDСolumnName    = "id"
	NameСolumnName  = "name"
	EmailСolumnName = "email"
	AgeСolumnName   = "age"
)
