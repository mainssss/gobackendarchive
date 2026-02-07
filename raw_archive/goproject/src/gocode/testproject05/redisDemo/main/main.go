package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis" //引入redis包
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
	_, err = conn.Do("set", "name", "tom猫猫")
	if err != nil {
		fmt.Println("set err:", err)
	}

	//3.通过go向redis读取数据string [key-val]
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("set err:", err)
		return
	}
	//因为返回的r是interface{}
	//因为name对应的值是string，因此我们需要转换
	//rname := r.(string)
	fmt.Println("get 完成,读取到:", r)
	fmt.Println("操作成功")
}
