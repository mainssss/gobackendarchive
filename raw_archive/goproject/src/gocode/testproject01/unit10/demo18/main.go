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
//中国人特有的方法:
func (person Chinese) NiuYangGe(){
	fmt.Println("东北文化-扭秧歌")
}

func (person American) disco(){
	fmt.Println("dicso")
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
	//断言:
	/*ch,flag := s.(Chinese)//看s是否能转成Chinese类型并且赋给ch变量,flag判断是否转成功 flag类型为bool
	if flag == true{
		ch.NiuYangGe()
	}else{
		fmt.Println("美国人不会扭秧歌")
	}*/
	
	/*if ch,flag := s.(Chinese);flag{
		ch.NiuYangGe()
	}else{
		fmt.Println("美国人不会扭秧歌")
	}*/

	switch s.(type){//type属于go语言中的一个关键字，固定写法:
	case Chinese:
		ch := s.(Chinese)
		ch.NiuYangGe()
	case American:
		am :=s.(American)
		am.disco()
	}
	
}


func main(){
	
	
	//创建一个中国人:
	c := Chinese{}
	//创建一个美国人:
	//a := American{}

	//美国人打招呼:
	//greet(a)
	//中国人打招呼:
	greet(c)

}