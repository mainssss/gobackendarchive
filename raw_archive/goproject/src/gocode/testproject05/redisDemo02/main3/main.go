package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("hmset", "Monster", "name", "snake", "age", 11, "skill", "bite")
	if err != nil {
		fmt.Println("hmset err:", err)
	} else {
		fmt.Println("hmset 完成...")
	}

	r, err := redis.Strings(conn.Do("hmget", "Monster", "name", "age", "skill"))
	if err != nil {
		fmt.Println("hmget err:", err)
	} else {
		fmt.Println("hmget 完成...")
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
