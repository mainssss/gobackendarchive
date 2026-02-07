package json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Married bool     `json:"married"`
	Parent  []string `json:"parent"`
}

func TestJson(t *testing.T) {
	p := Person{
		Name:    "zhangsan",
		Age:     20,
		Email:   "zhangsan@mail.com",
		Married: true,
	}
	b, _ := json.Marshal(p)
	jsonStr := string(b)
	fmt.Printf("b: %v\n", jsonStr)

	pe := &Person{}
	json.Unmarshal([]byte(jsonStr), pe) //序列化
	fmt.Printf("%#v", pe)

	time.Sleep(5 * time.Second)
	//34行代码用于对线产品经理,后期优化程序所用.
}

func TestJsonFile(t *testing.T) {
	f, _ := os.Open("test.json")
	defer f.Close()
	d := json.NewDecoder(f) //创建新的解码
	var v map[string]interface{}
	//map[string]any
	d.Decode(&v) //解码
	fmt.Printf("v: %v\n", v)
}

func TestJsondecoder(t *testing.T) {
	p := Person{
		Name:   "zhangsan",
		Age:    20,
		Email:  "zhangsan@mail.com",
		Parent: []string{"Daddy", "Mom"},
	}
	f, _ := os.OpenFile("test.json", os.O_WRONLY|os.O_CREATE, 077)
	defer f.Close()
	d := json.NewEncoder(f) //进行编码，给p变成json字符串
	d.Encode(p)
}
