package main

import (
	"fmt"
)

func main() {
	var num int
	num = 90
	//常量在声明的时候必须赋值
	const tax int = 0
	//常量不能修改
	//tax = 90
	fmt.Println(num, tax)
	//常量只能修饰bool、数值类型(int、float系列)、string类型

	/*
		const (
			a = 1
			b = 2
		)
		const (
			a = iota //给a赋值为0，在a的基础上依次+1
			b
			c
		)
		const(
			a = iota //0
			b = iota //1
			c,d = iota //2,2
		)
		const(

		)
	*/
}
