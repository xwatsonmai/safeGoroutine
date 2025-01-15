package goroutine

import (
	"context"
	"github.com/pkg/errors"
	"reflect"
)

//var poolCount = 99999999

type goRoutine struct {
	ctx          context.Context
	input        []any
	panicHandler func(err error)
}

func Go(function any, opts ...IGoRoutineOption) error {
	gr := &goRoutine{}
	for _, opt := range opts {
		opt.With(gr)
	}
	fnValue := reflect.ValueOf(function)
	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		return errors.New("第一个入参必须是函数")
	}
	if len(gr.input) != fnType.NumIn() {
		return errors.New("参数数量与函数的输入参数不匹配，请确认使用goroutine.WithInput{}传入参数")
	}
	in := make([]reflect.Value, len(gr.input))
	for i, arg := range gr.input {
		in[i] = reflect.ValueOf(arg)
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if gr.panicHandler != nil {
					gr.panicHandler(errors.New(err.(string)))
				}
			}
		}()
		fnValue.Call(in)
	}()
	return nil
}
