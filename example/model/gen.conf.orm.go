package model

import "github.com/ezbuy/redis-orm/orm"

//! conf.orm
type Unique interface {
	Key() string
	SQLFormat() string
	SQLParams() []interface{}
	SQLLimit() int
}

type Index interface {
	Unique
	Offset(n int)
	Limit(n int)
}

type Range interface {
	Index
	IncludeBegin(flag bool)
	IncludeEnd(flag bool)
}

type OrderBy interface {
	Index
	Ascend(flag bool)
}

type Finder interface {
	FindOne(unique Unique) (string, error)
	Find(index Index) ([]string, error)
	Range(scope Range) ([]string, error)
	OrderBy(sort OrderBy) ([]string, error)
}

type DBFetcher interface {
	FetchBySQL(sql string, args ...interface{}) ([]interface{}, error)
}

type ReferenceResult struct {
	db    Finder
	set   *orm.VSet
	times int
	err   error
}

func NewReferenceResult(db Finder) *ReferenceResult {
	return &ReferenceResult{
		db:  db,
		set: orm.NewVSet(),
	}
}

func (rr *ReferenceResult) DB(db Finder) *ReferenceResult {
	rr.db = db
	return rr
}

func (rr *ReferenceResult) Result() ([]string, error) {
	return rr.set.Values(rr.times), rr.err
}

func (rr *ReferenceResult) Values() []string {
	return rr.set.Values(rr.times)
}

func (rr *ReferenceResult) Unions() []string {
	return rr.set.Values(0)
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

func (rr *ReferenceResult) OrderBy(sort OrderBy) *ReferenceResult {
	rr.times = rr.times + 1
	if strs, err := rr.db.OrderBy(sort); err == nil {
		rr.set.SortAdd(1, strs...)
	} else {
		rr.err = err
	}
	return rr
}
