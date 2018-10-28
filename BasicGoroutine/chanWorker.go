package main

import (
	"fmt"
)

type responseMsg struct {
	hash int
}

type requestMsg struct {
	nouce        int
	responseChan chan *responseMsg
}

// requestChan： 只用来从channel中刚接收 *requestMsg 类型的数据
func chanWorker(requestChan chan *requestMsg) {
	// reqChanClosed := false
	// for {
	// 	select {
	// 	case req, ok := <-requestChan:
	// 		if ok {
	// 			respVal := req.nouce * 1234567
	// 			response := &responseMsg{respVal}
	// 			req.responseChan <- response
	// 		} else {
	// 			reqChanClosed = true
	// 		}
	// 	}
	// 	if reqChanClosed {
	// 		fmt.Println("chan worker closed")
	// 		break
	// 	}
	// }
	for req := range requestChan {
		respVal := req.nouce * 1234567
		response := &responseMsg{respVal}
		req.responseChan <- response
	}
}

func TestChanWorker() {
	fmt.Println("Go channels starting")

	requestChan := make(chan *requestMsg)

	defer close(requestChan)

	go chanWorker(requestChan)

	for i := 0; i < 5; i++ {
		request := &requestMsg{i, make(chan *responseMsg)}
		requestChan <- request
		response := <-request.responseChan
		fmt.Println("Got response:", response.hash)
	}

	// if _, ok := <-requestChan; !ok {
	// 	fmt.Println("channel closed")
	// }
}
