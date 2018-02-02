package model

import (
	"database/sql"
	"strings"
)

func IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}
