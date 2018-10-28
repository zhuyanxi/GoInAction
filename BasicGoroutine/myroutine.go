package main

import (
	"fmt"
	"time"
)

type myMsg struct {
	seqNum  int
	message string
}

func Test1() {
	fmt.Println("Go channels starting")

	ch1 := make(chan *myMsg)
	go sendChan(ch1, "channel 1")

	ch2 := make(chan *myMsg)
	go sendChan(ch2, "channel 2")

	// for msg := range ch {
	// 	fmt.Println("Msg ", msg.seqNum, ":", msg.message)
	// }
	ch1closed := false
	ch2closed := false
	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Println("channel 1:", msg.seqNum, " : ", msg.message)
			} else {
				if !ch1closed {
					fmt.Println("ch1 closed")
					ch1closed = true
				}
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Println("channel 2:", msg.seqNum, " : ", msg.message)
			} else {
				if !ch2closed {
					fmt.Println("ch2 closed")
					ch2closed = true
				}
			}
		}
		if ch1closed && ch2closed {
			break
		}
	}
}

// out: 只用来向channel中发送 *myMsg 类型的数据
func sendChan(out chan<- *myMsg, prefix string) {
	n := 0
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		out <- &myMsg{n, fmt.Sprintf("%s: %s", prefix, "moo")}
		n++
	}
	close(out)
}
