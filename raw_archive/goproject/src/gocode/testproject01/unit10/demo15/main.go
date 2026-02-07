package main
import "fmt"

//接口的定义:定义规则、定义规范，定义某种具体能力:
type sayHello interface{
	//声明没有实现的方法:
	sayHello()
}

//接口的实现:定义一个结构体:
//中国人:
type Chinese struct{
	name string
}

//实现接口的方法--->具体的实现：
func (person Chinese) sayHello(){
	fmt.Println("你好")
}

//接口的实现:定义一个结构体:
//美国人:
type American struct{
	name string
}

func (person American) sayHello(){
	fmt.Println("Hello")
}


//定义一个函数：专门用来各国人打招呼的函数，接收具备SayHello接口能力的变量：
func greet(s sayHello){//s:多态参数
	s.sayHello()
}

//自定义数据类型:
type integer int

func (i integer) sayHello(){
	fmt.Println("say hi +",i)
}
func main(){
	//定义一个SatHello类型的数组，里面存放American、Chinese结构体变量:
	var arr [3]sayHello
	arr[0] = American{"rose"}
	arr[1] = Chinese{"丽丽"}
	arr[2] = Chinese{"菲菲"}

	fmt.Println(arr)

	var i integer = 10
	var s sayHello = i

	s.sayHello() //say hi + 10
	//创建一个中国人:
	c := Chinese{}
	//创建一个美国人:
	a := American{}

	//美国人打招呼:
	greet(a)
	//中国人打招呼:
	greet(c)

	//直接用接口创建实例：出错:
	// var s sayHello
	// s.sayHello()

	// var s sayHello = c
	// s.sayHello()
}