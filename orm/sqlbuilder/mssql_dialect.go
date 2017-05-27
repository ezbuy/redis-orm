package sqlbuilder

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

var _ dbr.Dialect = MSSQLDialect{}

func mssqlQuoteIdent(s string) string {
	part := strings.SplitN(s, ".", 2)
	if len(part) == 2 {
		return mssqlQuoteIdent(part[0]) + "." + mssqlQuoteIdent(part[1])
	}
	return "[" + s + "]"
}

type MSSQLDialect struct{}

func (d MSSQLDialect) QuoteIdent(s string) string {
	return mssqlQuoteIdent(s)
}

func (d MSSQLDialect) EncodeString(s string) string {
	// http://www.postgresql.org/docs/9.2/static/sql-syntax-lexical.html
	return `N'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

func (d MSSQLDialect) EncodeBool(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func (d MSSQLDialect) EncodeTime(t time.Time) string {
	return `N'` + t.Format("2006-01-02 15:04:05.000") + `'`
}

func (d MSSQLDialect) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`'0x%x'`, b)
}

func (d MSSQLDialect) Placeholder(n int) string {
	return dialect.MySQL.Placeholder(n)
}
