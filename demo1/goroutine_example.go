package demo1

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sumN(N int) {
	defer wg.Done()
	sum := 0
	for i := 0; i <= N; i++ {
		sum += i
	}
	fmt.Printf("sum from 1 to %d sum is %d \n", N, sum)
}

func GoroutineTest() {
	wg.Add(1)
	go sumN(100)
	wg.Wait()
	fmt.Println("finished")
}
