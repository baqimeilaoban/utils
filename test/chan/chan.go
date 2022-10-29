package main

import (
	"fmt"
	"time"
)

func main() {
	testForrangeChan()
}

func testForrangeChan() {
	ch := make(chan int, 4)
	quitChan := make(chan bool)
	go func() {
		for i := range ch {
			println(i)
		}
		quitChan <- true
	}()
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("sleep 5 second")
			time.Sleep(time.Second * 5)
		}
		ch <- i
	}
	close(ch)  // 关闭通道，若是不关闭通道，会panic
	<-quitChan // 此通道用来通知是否ch已经完成
}
