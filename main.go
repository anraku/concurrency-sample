package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func expensiveMethod() error {
	for i := 0; i < 10000000; i++ {
	}
	return nil
}

func main() {
	start := time.Now()
	for i := 0; i < 30; i++ {
		expensiveMethod()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
