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

type WithInput []any

func (w WithInput) With(gr *goRoutine) {
	for _, v := range w {
		gr.input = append(gr.input, v)
	}
}

type WithPanicHandler func(err error)

func (w WithPanicHandler) With(gr *goRoutine) {
	gr.panicHandler = w
}

type WithWaitGroup struct {
	Wg *sync.WaitGroup
}

func (w WithWaitGroup) With(gr *goRoutine) {
	gr.wg = w.Wg
}
