package goroutine

import (
	"sync"
	"testing"
)

func TestGo(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		Go(func() {
			t.Log(v)
		}, WithWaitGroupDone{Wg: wg})
	}
	wg.Wait()
}

func TestGo2(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func() {
			t.Log(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
