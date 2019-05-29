package goroutine_channel

import "fmt"

//线程不安全的字典
func read(mp map[string]int) {
	fmt.Println(mp["fruit"])
}
func write(mp map[string]int) {
	mp["fruit"] = 2
}
func unsafeGo() {
	mp := make(map[string]int)
	go read(mp)
	write(mp)
}
