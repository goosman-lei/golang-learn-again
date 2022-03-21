package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		// Local redefine variable "i", the effect is same to parameter
		i := i
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Printf("ClosureVariable: %d Parameter: %d\n", i, v)
		}(i)

		// Adjust sleep time, we can see different result
		time.Sleep(time.Nanosecond * 200)
	}
	wg.Wait()
}