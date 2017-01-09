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

type MailboxPasswordOfUserUK struct {
	Mailbox  string
	Password string
}

func (u *MailboxPasswordOfUserUK) Key() string {
	strs := []string{
		"Mailbox",
		fmt.Sprint(u.Mailbox),
		"Password",
		fmt.Sprint(u.Password),
	}
	return fmt.Sprintf("unique:%s", strings.Join(strs, ":"))
}

func (u *MailboxPasswordOfUserUK) SQLFormat() string {
	conditions := []string{
		"mailbox = ?",
		"password = ?",
	}
	return strings.Join(conditions, " AND ")
}

func (u *MailboxPasswordOfUserUK) SQLParams() []interface{} {
	return []interface{}{
		u.Mailbox,
		u.Password,
	}
}

func (u *MailboxPasswordOfUserUK) SQLLimit() int {
	return 1
}

func (u *MailboxPasswordOfUserUK) Limit(n int) {
}

func (u *MailboxPasswordOfUserUK) Offset(n int) {
}

func (u *MailboxPasswordOfUserUK) UKRelation() UniqueRelation {
	return MailboxPasswordOfUserUKRelationRedisMgr()
}

//! indexes

type SexOfUserIDX struct {
	Sex    bool
	offset int
	limit  int
}

func (u *SexOfUserIDX) Key() string {
	strs := []string{
		"Sex",
		fmt.Sprint(u.Sex),
	}
	return fmt.Sprintf("index:%s", strings.Join(strs, ":"))
}

func (u *SexOfUserIDX) SQLFormat() string {
	conditions := []string{
		"sex = ?",
	}
	return fmt.Sprintf("%s %s", strings.Join(conditions, " AND "), orm.OffsetLimit(u.offset, u.limit))
}

func (u *SexOfUserIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Sex,
	}
}

func (u *SexOfUserIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *SexOfUserIDX) Limit(n int) {
	u.limit = n
}

func (u *SexOfUserIDX) Offset(n int) {
	u.offset = n
}

func (u *SexOfUserIDX) IDXRelation() IndexRelation {
	return SexOfUserIDXRelationRedisMgr()
}

//! ranges

type NameStatusOfUserRNG struct {
	Name         string
	StatusBegin  int32
	StatusEnd    int32
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
}

func (u *NameStatusOfUserRNG) Key() string {
	strs := []string{
		"Name",
		fmt.Sprint(u.Name),
		"Status",
	}
	return fmt.Sprintf("range:%s", strings.Join(strs, ":"))
}

func (u *NameStatusOfUserRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *NameStatusOfUserRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *NameStatusOfUserRNG) SQLFormat() string {
	conditions := []string{}
	conditions = append(conditions, "name = ?")
	conditions = append(conditions, fmt.Sprintf("status %s ?", u.beginOp()))
	conditions = append(conditions, fmt.Sprintf("status %s ?", u.endOp()))
	return fmt.Sprintf("%s %s", strings.Join(conditions, " AND "), orm.OffsetLimit(u.offset, u.limit))
}

func (u *NameStatusOfUserRNG) SQLParams() []interface{} {
	return []interface{}{
		u.Name,
		u.StatusBegin,
		u.StatusEnd,
	}
}

func (u *NameStatusOfUserRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *NameStatusOfUserRNG) Limit(n int) {
	u.limit = n
}

func (u *NameStatusOfUserRNG) Offset(n int) {
	u.offset = n
}

func (u *NameStatusOfUserRNG) Begin() int64 {
	return 0
}

func (u *NameStatusOfUserRNG) End() int64 {
	return 0
}

func (u *NameStatusOfUserRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *NameStatusOfUserRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *NameStatusOfUserRNG) ORDRelation() RangeRelation {
	return NameStatusOfUserRNGRelationRedisMgr()
}

//! orders
