package main
import "fmt"

// func   函数名（形参列表）（返回值类型列表）{
// 	执行语句..
// 	return+返回值列表
// } 

//自定义函数： 功能：两个数相加：
func cal (num1 int , num2 int) (int) {//如果返回值类型就一个的话，那么（）是可以省略不写的
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}



func main(){
	//求两数的和： 10+20
	//调用函数：
	sum := cal(10,20)
	fmt.Println(sum)
	// var num1 int = 10
	// var num2 int = 20
	//求和：
	// var sum int = 0
	// sum += num1
	// sum += num2
	// fmt.Println(sum)

	//功能：30 + 50
	var num3 int = 30
	var num4 int = 50
	//调用函数：
	sum1 := cal(num3,num4)
	fmt.Println(sum1)
	//求和：
	// var sum1 int = 0
	// sum1 += num3
	// sum1 += num4
	// fmt.Println(sum1)
}