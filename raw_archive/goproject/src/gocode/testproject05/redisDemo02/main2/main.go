package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入和读取数据
	//连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("redis连接失败:", err)
		return
	}
	defer conn.Close()

	//2.通过go向redis写入数据string [key-val]
	_, err = conn.Do("HMset", "animal2", "name", "汪汪", "age", 6)
	if err != nil {
		fmt.Println("hmset err:", err)
	} else {
		fmt.Println("Hset 完成")
	}

	//3.通过go向redis读取数据string [key-val]
	r, err := redis.Strings(conn.Do("Hmget", "animal2", "name", "age"))
	if err != nil {
		fmt.Println("hmget err:", err)
		return
	}

	//因为返回的r是interface{}
	//因为name对应的值是string，因此我们需要转换
	//rname := r.(string)
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
