package logrus

import (
	"context"
	"runtime"
	"time"
)

type ImmutableEntry struct {
	e *Entry
}

func (i *ImmutableEntry) Data() Fields {
	fields := Fields{}
	data := make(Fields, len(i.e.Data))
	for k, v := range i.e.Data {
		data[k] = v
	}
	return fields
}

func (i *ImmutableEntry) Time() time.Time {
	return i.e.Time
}

func (i *ImmutableEntry) Level() Level {
	return i.e.Level
}

func (i *ImmutableEntry) Caller() *runtime.Frame {
	return i.e.Caller
}

func (i *ImmutableEntry) Message() string {
	return i.e.Message
}

func (i *ImmutableEntry) Context() context.Context {
	return i.e.Context
}

func (i *ImmutableEntry) Err() string {
	return i.e.err
}
