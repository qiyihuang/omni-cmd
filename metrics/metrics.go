package metrics

import (
	"log"
	"time"
)

type Timer struct {
	name  string
	start int64
}

func StartTimer(name string) Timer {
	return Timer{name: name, start: time.Now().UnixMilli()}
}

func (t Timer) End() {
	log.Println(t.name, "took", time.Now().UnixMilli()-t.start, "ms to complete.")
}
