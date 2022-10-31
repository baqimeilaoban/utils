package main

import (
	"fmt"
	"time"
)

func main() {
	//testGoRoutine()
	goRoutineOpt()
}

/*
输出为：
get num is:  0
get num is:  1
get num is:  2
get num is:  3
get num is:  4
原因在于timer.After会生成一个新的channel，导致之前的channel被覆盖掉了，因此时间被重置了
*/
func testGoRoutine() {
	ch := make(chan int)
	// 开启一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("get num is: ", num)
			case <-time.After(2 * time.Second):
				fmt.Println("time is up")
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
}

/*
输出为：
get num is:  0
get num is:  1
time is up
get num is:  2
get num is:  3
get num is:  4
select具体的功能及用法如下：
1.select和case结合使用，每次执行select，都会只执行其中一个case或者执行default语句
2.当没有case或者default可以执行时，select则阻塞，等待直到一个case可以执行
3.当有多个case可以执行时，则随机选择一个case进行执行
4.case后面跟的必须是读或者写通道操作，否则编译出错
*/
func goRoutineOpt() {
	ch := make(chan int)
	idleDuration := 2 * time.Second
	idleDelay := time.NewTimer(idleDuration)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("get num is: ", num)
			case <-idleDelay.C:
				fmt.Println("time is up")
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
}
