package obgae

import (
	gae "appengine"
)

type ctxLogger struct {
	gae.Context
}

func newLogger(ctx gae.Context) (me *ctxLogger) {
	me = &ctxLogger{}
	me.Context = ctx
	return
}

func (me *ctxLogger) Fatal(err error) {
	me.Criticalf("FATAL: %+v", err)
	panic(err)
}