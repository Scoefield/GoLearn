package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("输入的参数有：", len(os.Args))
	//遍历os.Args切片，就得到所有命令行输入的参数
	for i, v := range os.Args {
		fmt.Printf("arg[%v]=%v\n", i, v)
	}
}
