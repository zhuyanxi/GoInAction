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
	go sendChan(ch)
	go sendChan(ch)

	for msg := range ch {
		fmt.Println("Msg ", msg.seqNum, ":", msg.message)
	}
}

func sendChan(out chan<- *myMsg) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		out <- &myMsg{i, "mooo"}
	}
	close(out)
}
