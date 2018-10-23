package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	//ListenOnce()
	//ListenForever()
	ListenUseGoroutine()
}

func ListenOnce() {
	ln, err := net.Listen("tcp", ":15440")
	if err != nil { // No ()s around cond
		fmt.Println("Error on listen: ", err)
		os.Exit(-1)
	}

	fmt.Println("Waiting for a connection via Accept(" + time.Now().String() + ")")
	conn, err := ln.Accept() // note:  Err doubly-declared.  Careful.
	if err != nil {
		fmt.Println("Error on accept: ", err)
		os.Exit(-1)
	}
	connNum := 0
	mconn := &myConn{
		conn:   conn,
		prefix: fmt.Sprintf("%d says", connNum),
	}
	handleConn(mconn)
	fmt.Println("Exiting")
}

func ListenForever() {
	ln, err := net.Listen("tcp", ":15440")
	if err != nil { // No ()s around cond
		fmt.Println("Error on listen: ", err)
		os.Exit(-1)
	}
	connNum := 0
	for {
		fmt.Println("Waiting for a connection via Accept(" + time.Now().String() + ")")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error on accept: ", err)
			os.Exit(-1)
		}
		mconn := &myConn{
			conn:   conn,
			prefix: fmt.Sprintf("%d says", connNum),
		}
		handleConn(mconn)
		connNum++
	}
	fmt.Println("Exiting")
}

func ListenUseGoroutine() {
	ln, err := net.Listen("tcp", ":15440")
	if err != nil { // No ()s around cond
		fmt.Println("Error on listen: ", err)
		os.Exit(-1)
	}
	connNum := 0
	for {
		fmt.Println("Waiting for a connection via Accept(" + time.Now().String() + ")")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error on accept: ", err)
			os.Exit(-1)
		}
		mconn := &myConn{
			conn:   conn,
			prefix: fmt.Sprintf("%d says", connNum),
		}
		go handleConn(mconn)
		connNum++
	}
	fmt.Println("Exiting")
}

// Go type inference is only partial:
// we had to know what type Accept returned
func handleConn(mconn *myConn) {
	fmt.Println("Reading once from connection")

	var buf [1024]byte
	n, err := mconn.conn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}

	fmt.Println(mconn.prefix, ":", string(buf[0:n]))
	mconn.conn.Close()
}

type myConn struct {
	conn   net.Conn
	prefix string
}
