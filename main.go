package main

import (
	"fmt"
	"github.com/xwatsonmai/safeGoroutine/goroutine"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	goroutine.Go(
		func() {
			test(1, 2)
		},
		goroutine.WithPanicHandler(func(err error) {
			fmt.Println("panic:", err)
		}),
		goroutine.WithWaitGroupDone{Wg: wg},
	)
	wg.Wait()
	//time.Sleep(time.Second * 51)
}

func test(a, b int) {
	fmt.Println(a, b)
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			panic("test")
		}
		time.Sleep(time.Second)
	}
}
