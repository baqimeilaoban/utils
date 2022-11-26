package main

import "fmt"

func main() {
	constValueNotChange()
}

// constValueNotChange 常量的值不能改变
func constValueNotChange() {
	fmt.Println(a)
	// a = "test1" 报错，cannot assign to a (untyped string constant "test")
}

const (
	a = "test"
)
