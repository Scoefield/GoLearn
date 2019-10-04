package main

import (
	"fmt"
	"reflect"
)

// 定义一个结构体
type Student struct {
	Name string
	Age  int
	Addr string
}

// 专门演示reflect
func reflectTest01(b interface{}) {
	// 1.获取reflect.TypeOf
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	// 2.获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v, rValue type:%T\n", rValue, rValue)
	// reflect.Value类型转int，再相加
	num1 := 20 + rValue.Int()
	fmt.Printf("num1=%v\n", num1)

	// 3.下面将rValue转回interface{}
	iV := rValue.Interface()
	// 将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

// 专门演示reflect对结构体的基本操作
func reflectTest02(b interface{}) {
	// 1.获取reflect.TypeOf
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	// 2.获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v, rValue type:%T\n", rValue, rValue)

	// 3.下面将rValue转回interface{}
	iV := rValue.Interface()
	fmt.Printf("iV=%v, iV type=%T\n", iV, iV)

	// 将 interface{} 通过断言转成需要的类型（这里断言为Student结构体类型）
	ret, ok := iV.(Student)
	if ok {
		fmt.Println("ret.name=", ret.Name)
	}
}

func main() {
	num := 100
	reflectTest01(num)

	stu := Student{"Jack", 23, "NewYork"}
	reflectTest02(stu)
}
