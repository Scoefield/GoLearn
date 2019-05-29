package goroutine_channel

import "fmt"

func channelOdd(ch chan<- int) {
	for i := 0; ; i++ {
		if i%2 == 1 {
			ch <- i
		}
	}
}
func chanelEven(ch chan<- int) {
	for i := 0; ; i++ {
		if i%2 == 0 {
			ch <- i
		}
	}
}
func channelSelect(chOne, chTwo <-chan int) {
	for {
		select {
		case a := <-chOne:
			fmt.Println("odd result: ", a)
		case b := <-chTwo:
			fmt.Println("even result: ", b)
		}
	}
}
func getChData() {
	chOne := make(chan int)
	chTwo := make(chan int)
	go channelOdd(chOne)
	go chanelEven(chTwo)
	go channelSelect(chOne, chTwo)
}
