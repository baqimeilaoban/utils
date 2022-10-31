package main

import "fmt"

func main() {
	//testAppendMap()
	testAppendMapOpt()
}

/*
输出结果为：
[map[name:bob] map[name:bob] map[name:bob]]
原因：
		user[`name`] = name
把之前的值覆盖了
*/
func testAppendMap() {
	var users []map[string]interface{}
	var user = make(map[string]interface{})
	var names = []string{
		"calvin",
		"jason",
		"bob",
	}
	for _, name := range names {
		user[`name`] = name
		users = append(users, user)
	}
	fmt.Println(users)
}

/*
输出结果：
[map[name:calvin] map[name:jason] map[name:bob]]
解决办法：在每一个循环中新声明一个局部变量
*/
func testAppendMapOpt() {
	var users []map[string]interface{}
	var names = []string{
		"calvin",
		"jason",
		"bob",
	}
	for _, name := range names {
		var user = make(map[string]interface{})
		user[`name`] = name
		users = append(users, user)
	}
	fmt.Println(users)
}
