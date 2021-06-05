package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFieldIN(t *testing.T) {
	in := NewFieldIN("test_field_in")
	in.Add(1)
	in.Add(2)
	in.Add(3)
	assert.Equal(t, "test_field_in IN (?,?,?)", in.SQLFormat())
	assert.Equal(t, []interface{}{1, 2, 3}, in.SQLParams())
}

func TestNewMultiFieldIN(t *testing.T) {
	in := NewMultiFieldIN([]string{"a", "b"})
	err := in.Add([]interface{}{1, 2})
	assert.NoError(t, err)
	err = in.Add([]interface{}{3, 4})
	assert.NoError(t, err)
	assert.Equal(t, "(a,b) IN ((?,?),(?,?))", in.SQLFormat())
	assert.Equal(t, []interface{}{1, 2, 3, 4}, in.SQLParams())
}
