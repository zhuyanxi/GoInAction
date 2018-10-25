package main

import (
	"fmt"
	"time"
)

type myMsg struct {
	seqNum  int
	message string
}

func main() {
	fmt.Println("Go channels starting")

	ch1 := make(chan *myMsg)
	go sendChan(ch1, "channel 1")

	ch2 := make(chan *myMsg)
	go sendChan(ch2, "channel 2")

	// for msg := range ch {
	// 	fmt.Println("Msg ", msg.seqNum, ":", msg.message)
	// }

	for {
		select {
		case msg := <-ch1:
			fmt.Println("channel 1:", msg.seqNum, " : ", msg.message)
		case msg := <-ch2:
			fmt.Println("channel 2:", msg.seqNum, " : ", msg.message)
		}
	}
}

func sendChan(out chan<- *myMsg, prefix string) {
	n := 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		out <- &myMsg{n, fmt.Sprintf("%s: %s", prefix, "moo")}
		n++
	}
	close(out)
}
