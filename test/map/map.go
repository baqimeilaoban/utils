package main

import "fmt"

func main() {
	pairs := [][2]string{
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
