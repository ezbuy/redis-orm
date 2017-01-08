package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/ezbuy/redis-orm/orm"
)

var (
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type UserBaseInfo struct {
	Id       int32  `db:"id"`
	Name     string `db:"name"`
	Mailbox  string `db:"mailbox"`
	Password string `db:"password"`
	Sex      bool   `db:"sex"`
}

//! object function

func (obj *UserBaseInfo) GetNameSpace() string {
	return "model"
}

func (obj *UserBaseInfo) GetClassName() string {
	return "UserBaseInfo"
}

func (obj *UserBaseInfo) GetTableName() string {
	return ""
}

func (obj *UserBaseInfo) GetColumns() []string {
	columns := []string{
		"`id`",
		"`name`",
		"`mailbox`",
		"`password`",
		"`sex`",
	}
	return columns
}

//! uniques

type MailboxPasswordOfUserBaseInfoUnique struct {
	Mailbox  string
	Password string
}

func (u *MailboxPasswordOfUserBaseInfoUnique) Key() string {
	strs := []string{
		"Mailbox",
		fmt.Sprint(u.Mailbox),
		"Password",
		fmt.Sprint(u.Password),
	}
	return fmt.Sprintf("unique:%s", strings.Join(strs, ":"))
}

func (u *MailboxPasswordOfUserBaseInfoUnique) SQLFormat() string {
	conditions := []string{
		"mailbox = ?",
		"password = ?",
	}
	return strings.Join(conditions, " AND ")
}

func (u *MailboxPasswordOfUserBaseInfoUnique) SQLParams() []interface{} {
	return []interface{}{
		u.Mailbox,
		u.Password,
	}
}

func (u *MailboxPasswordOfUserBaseInfoUnique) SQLLimit() int {
	return 1
}

//! indexes

type NameOfUserBaseInfoIndex struct {
	Name   string
	offset int
	limit  int
}

func (u *NameOfUserBaseInfoIndex) Key() string {
	strs := []string{
		"Name",
		fmt.Sprint(u.Name),
	}
	return fmt.Sprintf("index:%s", strings.Join(strs, ":"))
}

func (u *NameOfUserBaseInfoIndex) SQLFormat() string {
	conditions := []string{
		"name = ?",
	}
	return fmt.Sprintf("%s %s", strings.Join(conditions, " AND "), orm.OffsetLimit(u.offset, u.limit))
}

func (u *NameOfUserBaseInfoIndex) SQLParams() []interface{} {
	return []interface{}{
		u.Name,
	}
}

func (u *NameOfUserBaseInfoIndex) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *NameOfUserBaseInfoIndex) Limit(n int) {
	u.limit = n
}

func (u *NameOfUserBaseInfoIndex) Offset(n int) {
	u.offset = n
}

//! ranges

//! orders
