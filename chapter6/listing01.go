package main

import (
	"fmt"
	"runtime"
	"sync"
)

func L01() {
	// 分配逻辑处理器给调度器使用
	runtime.GOMAXPROCS(4)

	// wg用来等待程序完成
	// 计数加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// go func() {
	// 	//defer wg.Done()
	// 	//time.Sleep(time.Second)
	// 	fmt.Println("\nyet another goroutine")
	// }()

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 函数在退出时调用Done来通知main函数工作已经完成
		defer wg.Done()
		for count := 0; count < 6; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		defer wg.Done()
		for count := 0; count < 6; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
