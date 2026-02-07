package main

import (
	"fmt"
	"reflect"
)

//这里定义了一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_Age"`
	Score float32
	Sex   string
}

//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法，接收四个值,给monster赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	//获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取到a的类别
	kd := val.Kind()
	//判断如果传入的不是结构体，就退出
	if kd != reflect.Struct { //Struct是Kind包下的，Kind是reflect包下的
		fmt.Println("exnect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d: fields\n", num) //4
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d : 值为=%v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值(json的标签)
		tagVal := typ.Field(i).Tag.Get("json") //Type.Field和Value.Field是不一样的
		//如果该字段有标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d tag = %v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	//方法的排序默认是按照函数名的排序(ASCII码)
	val.Method(1).Call(nil) //获取到第二个方法，调用它

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value                   //声明了[]reflect.Value切片
	params = append(params, reflect.ValueOf(10)) //给params累加一个新的元素
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value，返回的还是切片
	fmt.Println("res=", res[0].Int()) //返回结果，返回的结果是[]reflect.Value
}
func main() {
	//创建了一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	//把Monster传给TestStruct函数
	TestStruct(a)
}
