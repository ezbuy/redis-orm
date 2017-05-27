package sqlbuilder

import "github.com/gocraft/dbr"

var _ Builder = &UpdateSet{}

type updateKV struct {
	k string
	v interface{}
}

type UpdateSet struct {
	kvs []updateKV
}

func (u *UpdateSet) Set(k string, v interface{}) *UpdateSet {
	u.kvs = append(u.kvs, updateKV{k: k, v: v})
	return u
}

func (u *UpdateSet) Build(d dbr.Dialect, buf dbr.Buffer) error {
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
