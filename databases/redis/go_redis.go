package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
)

var (
	network  = "tcp"
	address  = "127.0.0.1:6379"
	password = "123456"
)

// ************************** 获取连接 *******************************
func getConnect() redis.Conn {
	// 如果有密码才需要option
	option := redis.DialPassword(password)
	conn, err := redis.Dial(network, address, option)
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
	defer conn.Close()
	name, err := conn.Do("get", "name")
	if err != nil {
		log.Fatal("redis get error:", err)
	} else {
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
	defer conn.Close()
	res, err := redis.String(conn.Do("MGET", "name", "age"))
	if err != nil {
		log.Fatal("redis mget error:", err)
	} else {
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
	defer conn.Close()
	res, err := redis.String(conn.Do("LPOP", "list1"))
	if err != nil {
		log.Fatal("redis mget error:", err)
	} else {
		resType := reflect.TypeOf(res)
		fmt.Println("res_type: ", resType)
		fmt.Println("Lpop result: ", res)
		fmt.Println("result len: ", len(res))
	}
}

// ************************** 哈希操作 *******************************
func hashOpt() {
	conn := getConnect()
	defer conn.Close()
	_, err := conn.Do("HSET", "student", "name", "Jack", "age", 23)
	if err != nil {
		log.Fatal("HSET operation error:", err)
	}
	res, err := redis.Int(conn.Do("HGET", "student", "age"))
	if err != nil {
		log.Fatal("HGET operation is err: ", err)
	} else {
		resType := reflect.TypeOf(res)
		fmt.Println("res_type: ", resType)
		fmt.Println("Lpop result: ", res)
	}
}

// ************************** 管道操作（类似于并发） *******************************
/*Send：发送命令至缓冲区
Flush：清空缓冲区，将命令一次性发送至服务器
Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
*/
func pipineOpt() {
	conn := getConnect()
	defer conn.Close()
	conn.Send("HSET", "student", "name", "Jack", "age", 23)
	conn.Send("HSET", "student", "score", 103)
	conn.Send("HGET", "student", "score")
	conn.Flush()

	res, err := conn.Receive()
	if err != nil {
		log.Fatal("Receive error: ", err)
	} else {
		fmt.Println("Receive first result: ", res)
	}
	res2, err := conn.Receive()
	fmt.Println(res2)
	res3, err := conn.Receive()
	fmt.Println(res3)
}

// ************************** 连接池操作 *******************************
func connectPool() {
	Pool := redis.Pool{MaxIdle: 16, MaxActive: 32, IdleTimeout: 120, Dial: func() (conn redis.Conn, e error) {
		return redis.Dial(network, address)
	}}
	conn := Pool.Get()
	_, err := conn.Do("HSET", "student", "name", "Jack")
	if err != nil {
		log.Fatal("HSET pool error: ", err)
	}
	res, err := conn.Do("HGET", "student", "name")
	if err != nil {
		log.Fatal("HGET pool error: ", err)
	} else {
		fmt.Println(res)
	}
}
