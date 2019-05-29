package goroutine_channel

/*
defer 语句的应用场景 —— 释放锁。defer 语句总是要推迟到函数尾部运行，
所以如果函数逻辑运行时间比较长，这会导致锁持有的时间较长，
这时使用 defer 语句来释放锁未必是一个好注意。
*/
import (
	"fmt"
	"sync"
)

//让字典变的线程安全，就需要对字典的所有读写操作都使用互斥锁保护起来。
type SafeDict struct {
	data  map[string]int
	mutex *sync.Mutex
}

func NewSafeDict(mp map[string]int) *SafeDict {
	return &SafeDict{
		mp,
		&sync.Mutex{},
	}
}
func (d *SafeDict) Len() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return len(d.data)
}
func (d *SafeDict) Put(key string, value int) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	oldValue, ok := d.data[key]
	d.data[key] = value
	return oldValue, ok
}
func (d *SafeDict) Get(key string) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	oldValue, ok := d.data[key]
	return oldValue, ok
}
func (d *SafeDict) Delete(key string) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	oldValue, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return oldValue, ok
}

func Read(d *SafeDict) {
	fmt.Println(d.Get("apple"))
}
func Write(d *SafeDict) {
	d.Put("apple", 5)
}

func SafeGo() {
	d := NewSafeDict(map[string]int{
		"banana": 6,
		"pair":   8,
	})
	go Read(d)
	Write(d)
}
