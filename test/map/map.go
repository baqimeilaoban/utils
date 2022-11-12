package main

import "fmt"

func main() {
	mapSliceCompile()
}

// 输出ant
/*
正常情况下 map 字面量如果声明相同的 key 会编译错误，但是 pairs 第一层是切片不是数组，
编译器不会做编译时检查，因此会通过编译，然后字面量声明 map 会按照 key 顺序依次插入，
因此最后一次对 key “a” 的写入就是最终结果，而最后一次写入是 ant，最后的结果就是 ant。
*/
func mapSliceCompile() {
	pairs := [][]string{
		{"a", "apple"},
		{"a", "ant"},
		{"b", "bee"},
	}
	m := map[string]string{
		pairs[0][0]: pairs[0][1],
		pairs[1][0]: pairs[1][1],
		pairs[2][0]: pairs[2][1],
	}
	fmt.Println(m["a"])
}
