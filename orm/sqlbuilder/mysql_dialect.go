package sqlbuilder

import (
	"time"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

var _ dbr.Dialect = MySQLDialect{}

var mysqlTimeFormat = "2006-01-02 15:04:05.000000"

type MySQLDialect struct {
}

func (d MySQLDialect) QuoteIdent(s string) string {
	return dialect.MySQL.QuoteIdent(s)
}

func (d MySQLDialect) EncodeString(s string) string {
	return dialect.MySQL.EncodeString(s)
}

func (d MySQLDialect) EncodeBool(b bool) string {
	return dialect.MySQL.EncodeBool(b)
}

func (d MySQLDialect) EncodeTime(t time.Time) string {
	return `'` + t.Format(mysqlTimeFormat) + `'`
}

func (d MySQLDialect) EncodeBytes(b []byte) string {
	return dialect.MySQL.EncodeBytes(b)
}

func (d MySQLDialect) Placeholder(n int) string {
	return dialect.MySQL.Placeholder(n)
}
