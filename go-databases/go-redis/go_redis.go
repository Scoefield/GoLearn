package go_redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
)

var (
	network = "tcp"
	url = "127.0.0.1:6379"
	password = "123456"
)
// ************************** 获取连接 *******************************
func getConnect() redis.Conn {
	// 如果有密码才需要option
	option := redis.DialPassword(password)
	conn, err := redis.Dial(network, url, option)
	if err != nil {
		log.Fatal("get conn error:", err)
	}
	fmt.Println("connect success")
	return conn
}

// ************************** set和get操作 *******************************
func getAndSet() {
	conn := getConnect()
	_, err := conn.Do("set", "name", "Jack")
	if err != nil {
		log.Fatal("redis set error:", err)
	}
	name, err := conn.Do("get", "name")
	if err != nil {
		log.Fatal("redis get error:", err)
	}else {
		fmt.Println("Redis get name is: ", name)
	}
	// 设置key过期时间，10秒后过期
	//_, err = conn.Do("expire", "name", 10)
	//if err != nil {
	//	fmt.Println("key is expired")
	//}
}

// ************************** 批量Mset和Mget操作 *******************************
func manyGetSet() {
	conn := getConnect()
	_, err := conn.Do("MSET", "name", "Jack", "age", 22)
	if err != nil {
		log.Fatal("redis mset error:", err)
	}
	res, err := redis.String(conn.Do("MGET", "name", "age"))
	if err != nil {
		log.Fatal("redis mget error:", err)
	}else {
		resType := reflect.TypeOf(res)
		fmt.Println("res_type: ", resType)
		fmt.Println("Mget result: ", res)
		fmt.Println("result len: ", len(res))
	}
}

// ************************** 列表操作 *******************************
func listOpt() {
	conn := getConnect()
	_, err := conn.Do("LPUSH", "list1", "apple", "banana", "pair")
	if er != nil {
		log.Fatal("redis mset error:", err)
	}
	res, err := redis.String(conn.Do("LPOP", "list1"))
	if err != nil {
		log.Fatal("redis mget error:", err)
	}else {
		resType := reflect.TypeOf(res)
		fmt.Println("res_type: ", resType)
		fmt.Println("Lpop result: ", res)
		fmt.Println("result len: ", len(res))
	}
}


