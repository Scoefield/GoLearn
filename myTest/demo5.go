package main

import (
	"fmt"
	"sort"
)

type MyInterface interface {
	AInterface
	BInterface
	Say()
}

type AInterface interface {
	Hi()
}

type BInterface interface {
	Hello()
}

type integer int

func (i integer) Say() {
	fmt.Println("say it:", i)
}

func (i integer)Hello() {
	fmt.Println("hello")
}
func (i integer)Hi()  {
	fmt.Println("hi")
}

type Stu struct {
	Name string
	Age int
}

func (s *Stu)Hello() {
	fmt.Println("stu hello")
}

func main() {
	var i integer = 10
	var a MyInterface = i
	a.Say()

	var t interface{} = i
	fmt.Println(t)

	var stu = Stu{}
	var b BInterface = &stu
	b.Hello()

	var c = []int{7,3,5,6}
	var d = []float64{23.1, 33.4, 19.0}
	sort.Ints(c)
	fmt.Println(c)
	sort.Float64s(d)
	fmt.Println(d)

	var stuSlice []Stu
	s1 := Stu{"Jack", 23}
	s2 := Stu{"Mike", 24}
	stuSlice = append(stuSlice, s1, s2)
	fmt.Println(stuSlice)
}
