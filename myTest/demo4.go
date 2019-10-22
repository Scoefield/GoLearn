package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func (p *Person) GetName() string {
	return p.name
}
func (p *Person) GetAge() int {
	return p.age
}

func main() {
	p := Person{"Jack", 23}
	pType := reflect.TypeOf(p)
	fmt.Println(pType)

	pValue := reflect.ValueOf(p)
	fmt.Println(pValue)

	for i := 0; i < pValue.NumField(); i++ {
		fmt.Println(pValue.Field(i))
	}

}
