package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {

	f, _ := os.Create(fmt.Sprintf("trace_%s.out", time.Now().Format(time.RFC3339)))
	trace.Start(f)
	defer trace.Stop()

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber(&wg)
	go printNumber(&wg)

	wg.Wait()
}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Microsecond)
	}
}
