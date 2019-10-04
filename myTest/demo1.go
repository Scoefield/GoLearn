package main

import (
	"fmt"
	"strconv"
)

func main() {

	defer Recover()

	fmt.Println("hello go, this is a test")
	var i int = 4
	var j float32 = float32(i)
	fmt.Println(j)

	//var k int32 = 23
	//var l int64
	//l = int64(k) + 20

	var num int64 = 5
	ret := strconv.Itoa(int(num))
	fmt.Println(ret)

	var strBool string = "true"
	strRet, err := strconv.ParseBool(strBool)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", strRet, strRet)

	var str string = "12345"
	var inRet int64
	inRet, _ = strconv.ParseInt(str, 10, 64)
	fmt.Println(inRet)

}

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Recover:", r)
	}
}
