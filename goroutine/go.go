package goroutine

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

//var poolCount = 99999999

type goRoutine struct {
	ctx          context.Context
	input        []any
	panicHandler func(err error)
	wg           *sync.WaitGroup
}

func Go(function func(), opts ...IGoRoutineOption) error {
	gr := &goRoutine{}
	for _, opt := range opts {
		opt.With(gr)
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if gr.panicHandler != nil {
					gr.panicHandler(errors.New(err.(string)))
				}
			}
		}()
		if gr.wg != nil {
			defer gr.wg.Done()
		}
		function()
	}()
	return nil
}
