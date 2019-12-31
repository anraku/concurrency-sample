package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func expensiveMethod() error {
	time.Sleep(time.Second * 1)
	return nil
}

func main() {
	start := time.Now()
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			expensiveMethod()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
