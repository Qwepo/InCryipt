package types

import (
	"time"
)

type (
	PK   int64
	Unix int64
)

func (t *Unix) SetNow() {
	*t = Unix(time.Now().Unix())
}

func (t *Unix) Add(hour int) {
	*t = Unix(time.Now().Add(time.Duration(hour) * time.Hour).Unix())
}
