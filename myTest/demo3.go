package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Println(len(a), len(b), len(c), cap(a), cap(b), cap(c))

	e := []int{0}
	f := append(e[:0:0], e...)
	fmt.Println(reflect.TypeOf(f).Kind())
	fmt.Println(reflect.TypeOf(s).Kind())
}
