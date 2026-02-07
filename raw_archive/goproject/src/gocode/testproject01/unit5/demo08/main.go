package main

import "fmt"

//定义一个函数：
func test(num int) {
	fmt.Println(num)
}

//定义一个函数，把另一个函数作为形参：
func test02(num1 int, num2 float32, testFunc func(int)) {
	fmt.Println("---------test02")
}

type myFunc func(int)

func test03(num1 int, num2 float32, testFunc myFunc) {
	fmt.Println("---------test03")
}

//定义一个函数，求两个数的和，差
/*func test04(num1 int, num2 int) (int, int) {
	result01 := num1 + num2
	result02 := num1 - num2
	return result01, result02
}

//定义一个函数，求两个数的和，差
func test05(num1 int, num2 int) (sum int, sub int) {
	sub1 := num1 - num2
	sum1 := num1 + num2
	return sum1, sub1
}*/

func main() {
	//函数也是一种数据类型，可以赋值给一个变量
	a := test                                           //变量就是一个函数类型的变量
	fmt.Printf("a对应的类型是:%T,test函数对应的类型是:%T\n", a, test) //a对应的类型是:func(int),test函数对应的类型是:func(int)

	//通过该变量可以对函数进行调用
	a(10) //等价于 test(10)

	//调用test02对应的函数：
	test02(10, 3.19, test)
	test02(10, 3.19, a)

	//自定义数据类型：(相当于起别名):给int类型起别名，叫myInt类型
	type myInt int

	var num1 myInt = 30
	fmt.Println("num1", num1)

	var num2 int = 30
	num2 = int(num1) //虽然是别名，但是在Go中编译识别的时候还是认为myInt和int不是一种数据类型
	fmt.Println("num2", num2)

	test03(10, 9.8, a)
}
