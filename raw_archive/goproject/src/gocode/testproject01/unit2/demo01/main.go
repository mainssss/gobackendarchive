package main

import "fmt"

func main() {
	//1.变量的声明
	var age int
	//2.变量的赋值
	age = 18
	//3.变量的使用
	fmt.Println("age =", age)

	//变量声明和赋值可以合成一句
	var age2 int = 19
	fmt.Println("age2=", age2)

	/*变量的重复定义会报错：
	  	# command-line-arguments
	  	.\main.go:16:6: age redeclared in this block
	          previous declaration at .\main.go:6:6
	*/

	//不可以在赋值的时候给予是不匹配的类型
	var num int = 12
	fmt.Println("num=", num)
}
