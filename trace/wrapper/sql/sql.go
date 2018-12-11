package sql

import (
	"context"

	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
)

type Tracer struct {
	instance  string
	statement string
	dbtype    string
	user      string
	span      opentracing.Span
}

func NewMySQLTracer(ins string, user string) *Tracer {
	return &Tracer{
		dbtype:   "mysql",
		instance: ins,
		user:     user,
	}
}

func (t *Tracer) Do(ctx context.Context, statement string) {
	tracer := opentracing.GlobalTracer()
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}
	span = tracer.StartSpan("SQL Operation", opentracing.ChildOf(span.Context()))
	tags.DBInstance.Set(span, t.instance)
	tags.DBStatement.Set(span, statement)
	tags.DBType.Set(span, "mysql")
	tags.DBUser.Set(span, t.user)
	ctx = opentracing.ContextWithSpan(ctx, span)
	t.span = span
}

func (t *Tracer) Close() {
	t.span.Finish()
}
