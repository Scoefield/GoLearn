package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量，用于接受命令行的几个参数
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "localhost", "用户名")
	flag.IntVar(&port, "p", 3306, "端口号")

	// 这里有一个非常重要的操作，转换，必须要有
	flag.Parse()

	// 输出
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)

}
