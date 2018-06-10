package main

import (
	"math/rand"
	"sync"
	"time"
)

// wg用来等待程序完成
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//L01()
	//L04()
	L20()
}
