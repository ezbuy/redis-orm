package model

import "github.com/ezbuy/redis-orm/orm"

type SQL interface {
	SQLFormat() string
	SQLParams() []interface{}
	SQLLimit() int
	Offset(n int)
	Limit(n int)
}

//! conf.orm
type Unique interface {
	SQL
	Key() string
	UKRelation() UniqueRelation
}
type UniqueRelation interface {
	FindOne(key string) (string, error)
}

type Index interface {
	SQL
	Key() string
	IDXRelation() IndexRelation
}
type IndexRelation interface {
	Find(key string) ([]string, error)
}

type Range interface {
	SQL
	IncludeBegin(flag bool)
	IncludeEnd(flag bool)
	Begin() int64
	End() int64
	Revert(flag bool)
	Key() string
	RNGRelation() RangeRelation
}

type RangeRelation interface {
	Range(key string, start, end int64) ([]string, error)
	RevertRange(key string, start, end int64) ([]string, error)
}

type Finder interface {
	FindOne(unique Unique) (interface{}, error)
	Find(index Index) ([]interface{}, error)
	Range(scope Range) ([]interface{}, error)
	RevertRange(scope Range) ([]interface{}, error)
}

type DBFetcher interface {
	FetchBySQL(sql string, args ...interface{}) ([]interface{}, error)
}

type ReferenceResult struct {
	db     Finder
	set    *orm.VSet
	times  int
	offset int
	limit  int
	err    error
}

func NewReferenceResult(db Finder) *ReferenceResult {
	return &ReferenceResult{
		db:     db,
		set:    orm.NewVSet(),
		times:  0,
		offset: 0,
		limit:  -1,
	}
}

func (rr *ReferenceResult) DB(db Finder) *ReferenceResult {
	rr.db = db
	return rr
}

func (rr *ReferenceResult) Offset(n int) *ReferenceResult {
	rr.offset = n
	return rr
}

func (rr *ReferenceResult) Limit(n int) *ReferenceResult {
	rr.limit = n
	return rr
}

func (rr *ReferenceResult) Result() ([]interface{}, error) {
	return rr.set.Values(rr.times, rr.offset, rr.limit), rr.err
}

func (rr *ReferenceResult) Values() []interface{} {
	return rr.set.Values(rr.times, rr.offset, rr.limit)
}

func (rr *ReferenceResult) Unions() []interface{} {
	return rr.set.Values(0, rr.offset, rr.limit)
}

func (rr *ReferenceResult) Err() error {
	return rr.err
}

func (rr *ReferenceResult) FindOne(unique Unique) *ReferenceResult {
	rr.times = rr.times + 1
	if str, err := rr.db.FindOne(unique); err == nil {
		rr.set.Add(1, str)
	} else {
		rr.err = err
	}
	return rr
}

func (rr *ReferenceResult) Find(index Index) *ReferenceResult {
	rr.times = rr.times + 1
	if strs, err := rr.db.Find(index); err == nil {
		rr.set.Add(1, strs...)
	} else {
		rr.err = err
	}
	return rr
}

func (rr *ReferenceResult) Range(scope Range) *ReferenceResult {
	rr.times = rr.times + 1
	if strs, err := rr.db.Range(scope); err == nil {
		rr.set.SortAdd(1, strs...)
	} else {
		rr.err = err
	}
	return rr
}

func (rr *ReferenceResult) RevertRange(scope Range) *ReferenceResult {
	rr.times = rr.times + 1
	if strs, err := rr.db.RevertRange(scope); err == nil {
		rr.set.SortAdd(1, strs...)
	} else {
		rr.err = err
	}
	return rr
}
