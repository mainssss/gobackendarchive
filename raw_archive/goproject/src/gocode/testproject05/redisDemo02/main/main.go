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
	_, err = conn.Do("Hset", "animal1", "name", "汪汪")
	if err != nil {
		fmt.Println("hset err:", err)
	} else {
		fmt.Println("Hset 完成")
	}

	_, err = conn.Do("Hset", "animal1", "age", 10)
	if err != nil {
		fmt.Println("hset err:", err)
	} else {
		fmt.Println("Hset 完成")
	}

	//3.通过go向redis读取数据string [key-val]
	r1, err := redis.String(conn.Do("Hget", "animal1", "name"))
	if err != nil {
		fmt.Println("hget err:", err)
		return
	}

	r2, err := redis.Int(conn.Do("Hget", "animal1", "age"))
	if err != nil {
		fmt.Println("hget err:", err)
		return
	}
	//因为返回的r是interface{}
	//因为name对应的值是string，因此我们需要转换
	//rname := r.(string)
	fmt.Printf("操作成功 r1=%v\nr2=%v", r1, r2)
}
