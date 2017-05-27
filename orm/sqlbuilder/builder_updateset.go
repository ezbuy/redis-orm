package sqlbuilder

import "github.com/gocraft/dbr"

var _ Builder = &updateSet{}

type updateSetKV struct {
	k string
	v interface{}
}

func Set() *updateSet {
	return &updateSet{}
}

type updateSet struct {
	kvs []updateSetKV
}

func (u *updateSet) Add(k string, v interface{}) *updateSet {
	u.kvs = append(u.kvs, updateSetKV{k: k, v: v})
	return u
}

func (u *updateSet) Build(d dbr.Dialect, buf dbr.Buffer) error {
	i := 0
	for _, kv := range u.kvs {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(d.QuoteIdent(kv.k))
		buf.WriteString(" = ")
		buf.WriteString("?")

		buf.WriteValue(kv.v)
		i++
	}

	return nil
}
