package goroutine

import (
	"math/rand"
	"testing"
)

func BenchmarkGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Go(func() {
			testFunction(1000000)
		}, WithPanicHandler(func(err error) {
			// handle panic
		}))
	}
}

func BenchmarkGoKeyword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					// handle panic
				}
			}()
			testFunction(1000000)
		}()
	}
}

func testFunction(n int) {
	a := rand.New(rand.NewSource(99)).Float64()
	b := rand.New(rand.NewSource(99)).Float64()
	for i := 0; i < n; i++ {
		a += b
		a -= b
		a *= b
		if b != 0 {
			a /= b
		}
	}
}
