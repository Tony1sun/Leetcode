package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func main() {
// 	total := 0

// 	defer func() {
// 		time.Sleep(time.Second)
// 		fmt.Println("total: ", total)
// 		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
// 	}()

// 	var mutex sync.Mutex
// 	for i := 0; i < 2; i++ {
// 		go func() {
// 			mutex.Lock()
// 			defer mutex.Unlock()
// 			total += 1
// 		}()
// 	}
// }

func handle() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println("访问表1")
		wg.Done()
	}()

	go func() {
		fmt.Println("访问表2")
		wg.Done()
	}()

	go func() {
		fmt.Println("访问表3")
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	for i := 0; i < 4; i++ {
		go handle()

	}
	time.Sleep(time.Second)
}
