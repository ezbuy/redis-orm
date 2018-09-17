// Package cond provides orm-liked sql builder
package cond

import (
	"fmt"
	"strings"
)

const (
	maxQuestions        = 1 << 11
	defaultMaxQuestions = 1 << 8
)

// Cond defines sql conditions
type Cond struct {
	is   []int
	i32s []int32
	i64s []int64
	strs []string

	// ? series count
	questions uint16

	multiColumns MultiColumnsCond
}

// NewIntCond new Cond with int slice
func NewIntCond(ints []int) Cond {
	return Cond{
		is:   ints,
		i32s: []int32{},
		i64s: []int64{},
		strs: []string{},
	}
}

// NewInt32Cond new Cond with int32 slice
func NewInt32Cond(int32s []int32) Cond {
	return Cond{
		is:   []int{},
		i32s: int32s,
		i64s: []int64{},
		strs: []string{},
	}
}

// NewInt64Cond new Cond with int64 slice
func NewInt64Cond(int64s []int64) Cond {
	return Cond{
		is:   []int{},
		i32s: []int32{},
		i64s: int64s,
		strs: []string{},
	}
}

// NewStrCond new Cond with string slice
func NewStrCond(strs []string) Cond {
	return Cond{
		is:   []int{},
		i32s: []int32{},
		i64s: []int64{},
		strs: strs,
	}
}

// NewMultiColumns new Cond with multi columns search
func NewMultiColumns(columnsCond MultiColumnsCond) Cond {
	return Cond{
		is:           []int{},
		i32s:         []int32{},
		i64s:         []int64{},
		strs:         []string{},
		multiColumns: columnsCond,
	}
}

// NewDefaultMultiColumns new cond with default multi columns search
func NewDefaultMultiColumns(combineStr []string, key byte) Cond {
	return Cond{
		is:           []int{},
		i32s:         []int32{},
		i64s:         []int64{},
		strs:         []string{},
		multiColumns: NewMultiColumnsCond(string(key), combineStr),
	}
}

// SetMaxQuestion reset default MaxQuestion
func (c Cond) SetMaxQuestion(length uint16) error {
	if length > maxQuestions {
		return fmt.Errorf("query has too much questions,pls keep it below %d", maxQuestions)
	}
	c.questions = length
	return nil
}

// BuildPlaceholderStr builds raw ? holders str
func (c Cond) BuildPlaceholderStr() (holder string) {

	iLen := len(c.is)
	if iLen > 0 {
		return buildHolders(iLen)
	}

	i32Len := len(c.i32s)
	if i32Len > 0 {
		return buildHolders(i32Len)
	}

	i64Len := len(c.i64s)
	if i64Len > 0 {
		return buildHolders(i64Len)
	}

	strLen := len(c.strs)
	if strLen > 0 {
		return buildHolders(strLen)
	}

	// OTHER TYPES
	return c.multiColumns.BuildMultiColumnsPlaceholderStr()
}

// BuildPlaceholderStrWithQuote builds the (?,?...) str
func (c Cond) BuildPlaceholderStrWithQuote() (quoteHolder string) {
	hStr := c.BuildPlaceholderStr()
	if hStr == "" {
		return
	}
	return fmt.Sprintf("(%s)", hStr)
}

func buildHolders(count int) (holderStr string) {
	if count == 0 {
		return
	}
	holderStr = fmt.Sprintf("%s", strings.Repeat(",?", count)[1:])
	return
}

// BuildArgs builds args to redis-orm based structure ([]interface{})
func (c Cond) BuildArgs() (args []interface{}) {

	iLen := len(c.is)
	if iLen > 0 {
		for _, arg := range c.is {
			args = append(args, arg)
		}
	}

	i32Len := len(c.i32s)
	if i32Len > 0 {
		for _, arg := range c.i32s {
			args = append(args, arg)
		}
	}

	i64Len := len(c.i64s)
	if i64Len > 0 {
		for _, arg := range c.i64s {
			args = append(args, arg)
		}
	}

	strLen := len(c.strs)
	if strLen > 0 {
		for _, arg := range c.strs {
			args = append(args, arg)
		}
	}

	if len(args) > 0 {
		return args
	}

	// OTHER TYPES
	return c.multiColumns.BuildMultiColumnsArgs()
}
