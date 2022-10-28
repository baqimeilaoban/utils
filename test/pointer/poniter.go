package main

import "fmt"

func main() {
	var a int = 3
	var b *int = &a
	fmt.Println("a:", a)     // a: 3 值
	fmt.Println("b:", b)     // b: 0x1400001c090 地址
	fmt.Println("*b:", *b)   // *b: 3
	fmt.Println("&b:", &b)   // &b: 0x14000126018
	fmt.Println("*&a:", *&a) // *&a: 3
}
