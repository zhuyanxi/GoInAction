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

	ch := make(chan *myMsg)
	go sendChan(ch, "channel 1")
	go sendChan(ch, "channel 2")

	for msg := range ch {
		fmt.Println("Msg ", msg.seqNum, ":", msg.message)
	}
}

func sendChan(out chan<- *myMsg, prefix string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		out <- &myMsg{i, fmt.Sprintf("%s: %s", prefix, "moo")}
	}
	close(out)
}
