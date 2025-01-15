## This Go package encapsulates coroutines safely, preventing process crashes during concurrent use.
## 这个Go包安全地封装了协程，防止在并发使用时进程崩溃。

## 使用/Use

```go
package main

import (
	"fmt"
	"github.com/xwatsonmai/safeGoroutine/goroutine"

	"time"
)

func main() {
	goroutine.Go(
		test,
		goroutine.WithInput{1, 2}, // 传入参数 / input parameters
		goroutine.WithPanicHandler(func(err error) {
			fmt.Println("panic:", err)
		}), // panic处理 / panic handler
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

```

## 后续计划/Next Plan
- [ ] 增加Logger / Add Logger
- [ ] 增加可控制的协程并发数 / Add controllable coroutine concurrency
- [ ] 增加协程观察者 / Add coroutine observer