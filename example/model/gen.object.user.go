package model

import (
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"strings"
	"time"
)

var (
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type User struct {
	Id          int32     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Mailbox     string    `db:"mailbox" json:"mailbox"`
	Sex         bool      `db:"sex" json:"sex"`
	Longitude   float64   `db:"longitude" json:"longitude"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Description string    `db:"description" json:"description"`
	Password    string    `db:"password" json:"password"`
	HeadUrl     string    `db:"head_url" json:"head_url"`
	Status      int32     `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

//! object function

func (obj *User) GetNameSpace() string {
	return "model"
}

func (obj *User) GetClassName() string {
	return "User"
}

func (obj *User) GetTableName() string {
	return "users"
}

func (obj *User) GetColumns() []string {
	columns := []string{
		"`id`",
		"`name`",
		"`mailbox`",
		"`sex`",
		"`longitude`",
		"`latitude`",
		"`description`",
		"`password`",
		"`head_url`",
		"`status`",
		"`created_at`",
		"`updated_at`",
	}
	return columns
}
func (obj *User) GetIndexes() []string {
	idx := []string{
		"Sex",
	}
	return idx
}

func (obj *User) GetStoreType() string {
	return "hash"
}

func (obj *User) GetPrimaryName() string {
	return "Id"
}

//! uniques

type MailboxPasswordOfUserUnique struct {
	Mailbox  string
	Password string
}

func (u *MailboxPasswordOfUserUnique) Key() string {
	strs := []string{
		"Mailbox",
		fmt.Sprint(u.Mailbox),
		"Password",
		fmt.Sprint(u.Password),
	}
	return fmt.Sprintf("unique:%s", strings.Join(strs, ":"))
}

func (u *MailboxPasswordOfUserUnique) SQLFormat() string {
	conditions := []string{
		"mailbox = ?",
		"password = ?",
	}
	return strings.Join(conditions, " AND ")
}

func (u *MailboxPasswordOfUserUnique) SQLParams() []interface{} {
	return []interface{}{
		u.Mailbox,
		u.Password,
	}
}

func (u *MailboxPasswordOfUserUnique) SQLLimit() int {
	return 1
}

//! indexes

type SexOfUserIndex struct {
	Sex    bool
	offset int
	limit  int
}

func (u *SexOfUserIndex) Key() string {
	strs := []string{
		"Sex",
		fmt.Sprint(u.Sex),
	}
	return fmt.Sprintf("index:%s", strings.Join(strs, ":"))
}

func (u *SexOfUserIndex) SQLFormat() string {
	conditions := []string{
		"sex = ?",
	}
	return fmt.Sprintf("%s %s", strings.Join(conditions, " AND "), orm.OffsetLimit(u.offset, u.limit))
}

func (u *SexOfUserIndex) SQLParams() []interface{} {
	return []interface{}{
		u.Sex,
	}
}

func (u *SexOfUserIndex) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *SexOfUserIndex) Limit(n int) {
	u.limit = n
}

func (u *SexOfUserIndex) Offset(n int) {
	u.offset = n
}

//! ranges

//! orders
