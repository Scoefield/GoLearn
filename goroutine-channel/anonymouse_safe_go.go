package goroutine_channel

import "fmt"
import "sync"

type AnonSafeDict struct {
	data map[string]int
	*sync.Mutex
}

func AnonNewSafeDict(data map[string]int) *AnonSafeDict {
	return &AnonSafeDict{data, &sync.Mutex{}}
}

func (d *AnonSafeDict) AnonLen() int {
	d.Lock()
	defer d.Unlock()
	return len(d.data)
}

func (d *AnonSafeDict) AnonPut(key string, value int) (int, bool) {
	d.Lock()
	defer d.Unlock()
	oldValue, ok := d.data[key]
	d.data[key] = value
	return oldValue, ok
}

func (d *AnonSafeDict) AnonGet(key string) (int, bool) {
	d.Lock()
	defer d.Unlock()
	oldValue, ok := d.data[key]
	return oldValue, ok
}

func (d *AnonSafeDict) AnonDelete(key string) (int, bool) {
	d.Lock()
	defer d.Unlock()
	oldValue, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return oldValue, ok
}

func AnonWrite(d *AnonSafeDict) {
	d.AnonPut("banana", 5)
}

func AnonRead(d *AnonSafeDict) {
	fmt.Println(d.Get("banana"))
}

func AnonymouseMain() {
	d := AnonNewSafeDict(map[string]int{
		"apple": 2,
		"pear":  3,
	})
	go AnonRead(d)
	AnonWrite(d)
}
