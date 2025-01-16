package goroutine

import (
	"context"
	"sync"
)

type IGoRoutineOption interface {
	With(*goRoutine)
}

type WithContext struct {
	Ctx context.Context
}

func (w WithContext) With(gr *goRoutine) {
	gr.ctx = w.Ctx
}

type WithPanicHandler func(err error)

func (w WithPanicHandler) With(gr *goRoutine) {
	gr.panicHandler = w
}

type WithWaitGroupDone struct {
	Wg *sync.WaitGroup
}

func (w WithWaitGroupDone) With(gr *goRoutine) {
	gr.wg = w.Wg
}
