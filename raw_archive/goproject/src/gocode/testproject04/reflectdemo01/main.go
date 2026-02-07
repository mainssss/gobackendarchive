package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {

	//通过反射获取传入变量的type,kind,值(Value):
	//1.先获取到reflect.type
	rtyp := reflect.TypeOf(b) //这个为反射类型的type
	fmt.Println("rtyp=", rtyp)

	//2.获取到reflect.Value类型
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	//n3 := rVal.Float()  //传进来的值是int类型，转换后的是常量，不能改变
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//3.将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

//专门演示对结构体的反射
func reflectTest02(b interface{}) {
	//通过反射获取传入变量的type,kind,值(Value):
	//1.先获取到reflect.type
	rtyp := reflect.TypeOf(b) //这个为反射类型的type
	fmt.Println("rtyp=", rtyp)

	//2.获取到reflect.Value类型
	rVal := reflect.ValueOf(b)

	//3.获取变量对应的kind
	//(1)rVal.kind() ==>
	kind1 := rVal.Kind()
	//(2)rTyp.kind() ==>
	kind2 := rtyp.Kind()
	fmt.Printf("kind1=%v kind2=%v \n", kind1, kind2)

	//将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV=%v iV type=%T\n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//简单的使用了带检测的类型断言
	//同学们可以使用swith的方法
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {

	//编写一个案例,
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1.先定义一个int
	var num int = 100
	reflectTest01(num)

	//定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
