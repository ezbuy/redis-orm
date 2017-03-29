package model

type SQL interface {
	SQLFormat(limit bool) string
	SQLParams() []interface{}
	SQLLimit() int
	Offset(n int)
	Limit(n int)
}

//! conf.orm
type PrimaryKey interface {
	Key() string
	SQLFormat() string
	SQLParams() []interface{}
	Columns() []string
	Parse(key string) error
}

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
	PositionOffsetLimit(len int) (int, int)
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
	PositionOffsetLimit(len int) (int, int)
	RNGRelation() RangeRelation
}

type RangeRelation interface {
	Range(key string, start, end int64) ([]string, error)
	RangeRevert(key string, start, end int64) ([]string, error)
}

type Finder interface {
	FindOne(unique Unique) (PrimaryKey, error)
	Find(index Index) (int64, []PrimaryKey, error)
	Range(scope Range) (int64, []PrimaryKey, error)
	RangeRevert(scope Range) (int64, []PrimaryKey, error)
}

type DBFetcher interface {
	FetchBySQL(sql string, args ...interface{}) ([]interface{}, error)
}
