package main

import (
	"fmt"
	"regexp"
)

func main() {
	//正则匹配数字
	if m, _ := regexp.MatchString("^[0-9]+$", "12f345"); !m {
		fmt.Println("false")
	}
	//匹配中文
	if m, _ := regexp.MatchString("^\\p{Han}+$", "测试匹配汉子"); !m {
		fmt.Println("false")
	}
}
