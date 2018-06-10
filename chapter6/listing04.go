package main

import (
	"fmt"
	"runtime"
)

func L04() {

	// 分配逻辑处理器给调度器使用
	runtime.GOMAXPROCS(3)
	//runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Complete", prefix)
}
