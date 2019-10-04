package main

import "fmt"

// 传统测试方法
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	ret := addUpper(10)
	if ret == 55 {
		fmt.Println("test success")
	} else {
		fmt.Println("test fail")
	}
}
