package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()
	panic("main panic")
}
