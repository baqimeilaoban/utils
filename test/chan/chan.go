package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//testForrangeChan()
	//testChannelComplete()
	testChanSelect()
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

func testChannelComplete() {
	rNum := 100000
	r := make(chan int, rNum)
	var wg sync.WaitGroup
	for i := 0; i < rNum; i++ {
		wg.Add(1)
		go func(x int, r chan int) {
			defer wg.Done()
			if x%3 == 0 && x%23 == 0 {
				r <- x
			}
		}(i, r)
	}
	wg.Wait()
	close(r)
	for i := range r {
		fmt.Println(i)
	}
}

func goRoutineD(ch chan int, i int) {
	time.Sleep(time.Second * 3)
	ch <- i
	close(ch) // 关闭通道
}
func goRoutineE(chs chan string, i string) {
	time.Sleep(time.Second * 3)
	chs <- i
	close(chs) // 关闭通道
}

func testChanSelect() {
	ch := make(chan int, 5)
	chs := make(chan string, 5)

	go goRoutineD(ch, 5)
	go goRoutineE(chs, "ok")
	exitChFlag := false
	exitChsFlag := false
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				ch = make(chan int)
				fmt.Println("ch通道已关闭")
				exitChFlag = true
			} else {
				fmt.Println(" received the data ", msg)
			}
		case msgs, ok := <-chs:
			if !ok {
				chs = make(chan string)
				fmt.Println("chs通道已关闭")
				exitChsFlag = true
			} else {
				fmt.Println(" received the data ", msgs)
			}
		default:
			fmt.Println("channel中没有数据")
			time.Sleep(1 * time.Second)
		}
		if exitChFlag && exitChsFlag {
			fmt.Println("跳出循环")
			break
		}
	}

}
