package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	for i := 0; i < 5; i++ {
		s := "第" + strconv.Itoa(i) + "个"
		Client([]byte(s))
	}
	//time.Sleep(1000)
}
func Client(sms []byte) {
	//for {
	//sms := make([]byte, 128)
	conn, err := net.Dial("tcp", ":15440")
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	//fmt.Print("请输入要发送的消息:")
	// _, err = fmt.Scan(&sms)
	// if err != nil {
	// 	fmt.Println("数据输入异常:", err.Error())
	// }
	conn.Write(sms)

	// buf := make([]byte, 128)
	// c, err := conn.Read(buf)
	// if err != nil {
	// 	fmt.Println("读取服务器数据异常:", err.Error())
	// }
	// fmt.Println(string(buf[0:c]))

	conn.Close()
	//}
}
