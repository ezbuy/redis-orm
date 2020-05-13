package orm

import (
	"database/sql"
	"strings"
)

// IsErrNotFound represents the mysql db rows not found error
// should be compatible with FetchByPrimaryKey and FetchByPK
func IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}
