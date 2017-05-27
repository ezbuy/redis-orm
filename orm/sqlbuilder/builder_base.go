package sqlbuilder

import "github.com/gocraft/dbr"

func Expr(query string, value ...interface{}) dbr.Builder {
	return dbr.Expr(query, value...)
}

func And(cond ...dbr.Builder) dbr.Builder {
	if len(cond) == 1 {
		return cond[0]
	}

	return dbr.And(cond...)
}

func Or(cond ...dbr.Builder) dbr.Builder {
	if len(cond) == 1 {
		return cond[0]
	}

	return dbr.Or(cond...)
}

func Eq(column string, value interface{}) dbr.Builder {
	return dbr.Eq(column, value)
}

func Neq(column string, value interface{}) dbr.Builder {
	return dbr.Neq(column, value)
}

func Gt(column string, value interface{}) dbr.Builder {
	return dbr.Gt(column, value)
}

func Gte(column string, value interface{}) dbr.Builder {
	return dbr.Gte(column, value)
}

func Lt(column string, value interface{}) dbr.Builder {
	return dbr.Lt(column, value)
}

func Lte(column string, value interface{}) dbr.Builder {
	return dbr.Lte(column, value)
}

func I(s string) dbr.I {
	return dbr.I(s)
}

// TODO: pagination, orderby
