package main

import (
	"fmt"
	"sync"

	"sync/atomic"
)

// atomic CAS 原子操作
func main() {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 失败一直重试，以保证count自增10000次
			for {
				old := atomic.LoadInt64(&count)
				if atomic.CompareAndSwapInt64(&count, old, old+1) {
					break
				}
				fmt.Println("swap failed, current old: ", old)
			}

		}(&wg)
	}
	wg.Wait()

	// count = 10000
	fmt.Println("count = ", count)
}
