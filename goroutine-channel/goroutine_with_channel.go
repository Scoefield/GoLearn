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

/********************** channel阻塞情况，及其解决方法 *********************/
// 阻塞死锁的情况
// 这里的运行结果会报错：fatal error: all goroutines are asleep - deadlock!
func dealLock() {
	ch := make(chan int)
	ch <- 1
	val := <-ch
	fmt.Println(val)
}

// 这里的运行结果为：1  不会报错，通过初试化channel时，让channel容量大于0，是异步非阻塞，解决了死锁问题
func chWithCap() {
	ch := make(chan int, 1)
	ch <- 1
	val := <-ch
	fmt.Println(val)
}

/*
首先，管道的容量依然为0；但是，通过go关键字来创建一个协程，这这个单独的协程中就可以做好准备，
向管道中写入这个值，并且也不会阻塞主线程；在主线程中，消费者做好准备从管道中读取值；
在某个时刻，生产者和消费者都准备好了，进行通信，这就不会导致死锁了。
*/
func chWithRoutine() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	val := <-ch
	fmt.Println(val)
}

/********************** channel方向，close, for *********************/
/*
1、sender函数中，持续地往管道中写入int类型的消息，并且在i为5的时候调用close手动关闭了管道，并且跳出了循环。
	这里需要注意的是，不能再向已经关闭的管道中写入值，因此如果没有上面的break，会触发panic
2、receiver函数中，使用for-range的形式从管道中读取值，在管道被关闭之后，会自动的结束循环
3、同时，我们还注意到，sender函数的形参类型是 chan<- int，receiver函数的形参类型是 <-chan int，
	这代表着管道是单向的，分别只能向管道中写入消息、读取消息
*/
func sendToCh(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		if i == 5 {
			close(ch)
			break
		}
	}
}
func receiveCh(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
func chMain() {
	ch := make(chan int)
	go sendToCh(ch)
	receiveCh(ch)
}
