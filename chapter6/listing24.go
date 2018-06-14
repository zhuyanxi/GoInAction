package main

import (
	"fmt"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

func L24() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 0; gr < numberGoroutines; gr++ {
		go work(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 0; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
		fmt.Printf("*****************************\nTask %d Added to Channel\n*****************************\n", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有goroutine 退出
	close(tasks)
	// 等待所有工作完成
	wg.Wait()
}

// work作为goroutine 启动来处理
// 从有缓冲的通道传入的工作
func work(tasks chan string, worker int) {
	defer wg.Done()
	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down the channel %s.\n", worker, <-tasks)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker : %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		//sleep := rand.Int63n(100)
		time.Sleep(time.Duration(2000) * time.Millisecond)
		fmt.Println("==============================================================")

		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
