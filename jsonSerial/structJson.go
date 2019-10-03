package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Sal   float64 `json:"sal"`
	Skill string  `json:"skill"`
}

// 测试结构体序列化
func testStruct() {
	// 演示
	monster := Monster{
		Name:  "Jack",
		Age:   23,
		Sal:   13000.0,
		Skill: "Coding",
	}
	// 将monster序列化
	jsonStruct, err := json.Marshal(monster)
	if err != nil {
		panic(err)
	}
	// 输出序列化后的struct结果
	fmt.Printf("struct序列化后的结果为：%v\n", string(jsonStruct))
}

// 测试map序列化
func testMap() {
	// 定义一个map用于测试
	student := make(map[string]interface{})
	student["name"] = "Mike"
	student["age"] = 32
	student["address"] = "NewYork"
	// 将student序列化
	jsonMap, err := json.Marshal(&student)
	if err != nil {
		panic(err)
	}
	// 输出序列化map后的结果
	fmt.Printf("map序列化后的结果为：%v\n", string(jsonMap))

}

// 演示切片slice序列化
func testSlice() {
	// 定义一个map类型的切片，用于存放数据
	var sliceRet []map[string]interface{}

	m := make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	// 将map1结果加入到切片中
	sliceRet = append(sliceRet, m)

	m2 := make(map[string]interface{})
	m2["username"] = "user1"
	m2["age"] = 18
	m2["sex"] = "man"
	// 将map2结果加入到切片中
	sliceRet = append(sliceRet, m2)

	// 序列化切片slice
	data, err := json.Marshal(&sliceRet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("slice序列化后的结果为：%v\n", string(data))
}

func main() {
	// 演示将struct序列化
	testStruct()
	// 演示将map序列化
	testMap()
	// 演示将slice序列化
	testSlice()
}
