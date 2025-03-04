package goroutine

import (
	"context"
	"fmt"
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
				errData := fmt.Errorf("%v", err)
				if gr.panicHandler != nil {
					gr.panicHandler(errData)
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
