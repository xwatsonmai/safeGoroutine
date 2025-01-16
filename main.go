package main

import (
	"fmt"
	"github.com/xwatsonmai/safeGoroutine/goroutine"
	"time"
)

func main() {
	goroutine.Go(
		func() {
			test(1, 2)
		},
		goroutine.WithPanicHandler(func(err error) {
			fmt.Println("panic:", err)
		}),
	)
	time.Sleep(time.Second * 51)
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
