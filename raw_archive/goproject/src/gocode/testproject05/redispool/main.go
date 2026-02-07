package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//初始化pool
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //可用连接
		MaxActive:   0,   //可用连接数
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接的代码，连接哪个IP
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//先从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "jack")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	fmt.Println("r:", r)
	//conn.Close()
	//如果连接池关闭，就放/取不到东西了
}
