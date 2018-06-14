package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg用来等待程序完成
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func testDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func main() {
	//testDefer()

	//L01()
	//L04()
	//L20()
	//L22()
	L24()

	// ba := make(chan int)
	// wg.Add(1)
	// go func() {
	// 	fmt.Println(<-ba)
	// 	wg.Done()
	// }()
	// ba <- 3
	// wg.Wait()
	// close(ba)
}
