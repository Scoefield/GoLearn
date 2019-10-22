package main

import (
	"GitCode/GoLearn/jsons"
	"fmt"
)

type user struct {
	Name   string
	Blog   string
	Wechat string
}

func main() {
	u := user{"飞雪无情", "http://www.flysnow.org/", "flysnow_org"}
	b, err := jsons.MyMarshalIndent(u, "", " ")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}
