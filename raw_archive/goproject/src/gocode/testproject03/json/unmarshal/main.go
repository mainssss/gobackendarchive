package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

//演示反序列化将json字符串反序列化成结构体
func unmarshalStruct() {
	//说明str 在项目开发中是通过网络传输获取的..或者是读取文件获取到的
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011/11/11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	//定义一个monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.name=%v\n", monster, monster.Name)
}

//演示反序列化将json字符串反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"

	//定义一个map
	var a map[string]interface{}

	//反序列化会自动make序列化空间,因为make操作被封装到Unmarshal函数中去了
	//反序列化
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v \n", a)
}

//演示将json串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西 哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	//定义一个Slice
	var slice []map[string]interface{}
	//反序列化,不需要make，因为make被封装到Umarshal函数了
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v \n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
