package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Stu struct {
	Name  string  `json:"name"`
	Grade float32 `json:"grade"`
}

func TestStu(name string, fs string) (err error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	defer conn.Close()

	_, err = conn.Do("hset", "stu", name)
	if err != nil {
		fmt.Println("存入失败")
		return
	}

	return
}

//var decide bool = false
var key int
var name string
var fs string

func main() {
	for {
		fmt.Println("---欢迎使用学生分数管理系统---")
		fmt.Println("1.查询学生分数")
		fmt.Println("2.录入学生分数")
		fmt.Println("3.删除学生分数")
		fmt.Println("4.退出")
		fmt.Println("请输入你的选择:")
		fmt.Scanf("\n%d", &key)

		switch key {
		case 1:
			fmt.Println("请输入学生姓名:")
			fmt.Scanf("%s\n", &name)
			fmt.Println("请输入学生分数:")
			fmt.Scanf("%s\n", &fs)

			go TestStu(name, fs)
		case 2:
			fmt.Println()
		case 3:
			fmt.Println("删除")
		case 4:
			break

		default:
			fmt.Println("你的输入有误")
		}
	}
}
