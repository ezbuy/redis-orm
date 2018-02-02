package model

import "strings"

func IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found")
}
