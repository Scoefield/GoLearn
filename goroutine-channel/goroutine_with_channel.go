package goroutine_channel

/*
多线程运行速度明显比单线程快，测试对比结果：
=== RUN   TestSerial
sum = 7999999998000000000, spend time: 1.7382813s --- PASS: TestSerial (1.74s)
=== RUN   TestParallel
sum = 7999999998000000000, spend time: 1.1835937s--- PASS: TestParallel (1.18s)
PASS
ok      _/D_/Do/GitCode/GoLearn/goroutine-channel       2.979s
*/
import (
	"fmt"
	"time"
)

// 单线程
func serial() {
	sum, count, start := 0, 0, time.Now()
	for count < 4e9 {
		sum += count
		count += 1
	}
	end := time.Now()
	fmt.Printf("sum = %d, spend time: %s ", sum, end.Sub(start))
}

//多线程性能明显提升
func parallel() {
	sum, count, start, ch := 0, 0, time.Now(), make(chan int, 4)
	for count < 4 {
		go func(count int) {
			value := 0
			for i := count * 1e9; i < (count+1)*1e9; i++ {
				value += i
			}
			ch <- value
		}(count)
		count++
	}
	for count > 0 {
		sum += <-ch
		count--
	}
	end := time.Now()
	fmt.Printf("sum = %d, spend time: %s", sum, end.Sub(start))
}
