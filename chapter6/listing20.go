package main

import (
	"fmt"
	"math/rand"
)

func L20() {
	court := make(chan int)
	wg.Add(2)

	//启动两个选手
	go player("Nadal", court)
	go player("Zero", court)

	//发球
	court <- 1

	//等待游戏结束
	wg.Wait()
}

// player 模拟一个选手打网球
func player(name string, court chan int) {
	defer wg.Done()
	for {
		//等待球被击打回来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		//选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed, %d\n", name, n)

			//通道关闭，表示我们输了
			close(court)
			return
		}

		//显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		//将球打向球手
		court <- ball
	}
}
