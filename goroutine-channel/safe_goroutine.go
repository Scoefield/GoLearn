package goroutine_channel

import (
	"fmt"
	"sync"
)

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
