package unitestApplication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 给Monster绑定Store方法，实现序列化monster写到文件中去
func (m *Monster) Store() bool {
	// 先序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal error:", err)
		return false
	}
	// 写入文件
	filePath := "monster.json"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file error:", err)
		return false
	}
	return true
}

// 给Monster绑定ReStore方法，实现可以将一个序列化的monster，从文件中读取
func (m *Monster) ReStore() bool {
	// 文件路径
	filePath := "monster.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file error:", err)
		return false
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return false
	}
	return true
}
