package sqlbuilder

import (
	"github.com/gocraft/dbr"
)

type Builder dbr.Builder

var (
	MySQL = SQLBuilder{
		d: MySQLDialect{},
	}

	MSSQL = SQLBuilder{
		d: MSSQLDialect{},
	}
)

type SQLBuilder struct {
	d dbr.Dialect
}

func (s *SQLBuilder) Build(b Builder) (string, error) {
	return dbr.InterpolateForDialect("?", []interface{}{b}, s.d)
}

func (s *SQLBuilder) MustBuild(b Builder) string {
	out, err := s.Build(b)
	if err != nil {
		panic(err)
	}

	return out
}
