package main

import "fmt"

func main() {
	testDeferRegister()
}

/*
输出结果：
is true
defer1
原因：defer需要在return之前注册后，才能按照后人先出的出栈方式出栈
*/
func testDeferRegister() {
	a := true
	defer func() {
		fmt.Println("defer1")
	}()
	if a == true {
		fmt.Println("is true")
		return
	}
	defer func() {
		fmt.Println("defer2")
	}()
}
