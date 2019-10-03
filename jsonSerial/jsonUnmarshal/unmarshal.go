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

// 演示将json字符串，反序列化为struct
func unmarshalStruct() {
	// 定义一个struct用于接收反序列化的结果
	monster := Monster{}

	// json字符串
	str := `{"name":"Jack","age":23,"sal":13000,"skill":"Coding"}`

	// 进行反序列化
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		panic(err)
	}
	fmt.Printf("反序列化后的结构体结果为：%v\n", monster)
	fmt.Printf("monster.name=%v, monster.age=%v\n", monster.Name, monster.Age)
}

// 演示将json字符串，反序列化为map
func unmarshalMap() {
	// 定义一个map变量
	//retMap := make(map[string]interface{})
	var retMap map[string]interface{}

	// 定义一个map json字符串变量
	str := `{"address":"NewYork","age":32,"name":"Mike","skill":["coding","dancing"]}`

	// 进行反序列化，注意：反序列化map时，不需要make操作了，因为make操作被封装到了unmarshal函数里面去了
	err := json.Unmarshal([]byte(str), &retMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("反序列化的map结果为：%v\n", retMap)
}

// 演示将json字符串，反序列化为slice
func unmarshalSlice() {
	str := `[{"age":29,"sex":"female","username":"user2"},{"age":18,"sex":"man","username":"user1"}]`

	// 定义一个slice
	var retSlice []map[string]interface{}
	// 进行反序列化，注意：同理map，反序列化slice时，不需要make操作了，因为make操作被封装到了unmarshal函数里面去了
	err := json.Unmarshal([]byte(str), &retSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("反序列化后slice的结果为：%v\n", retSlice)

}

func main() {
	// 演示反序列化的结构体
	unmarshalStruct()
	// 演示反序列化后的map
	unmarshalMap()
	// 演示反序列化后的slice
	unmarshalSlice()
}
