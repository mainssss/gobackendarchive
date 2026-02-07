package main
import "fmt"



//自定义函数： 功能：两个数相加：
func cal (num1 int , num2 int) int {//如果返回值类型就一个的话，那么（）是可以省略不写的
	var sum int = 0
	sum += num1
	sum += num2
	return sum
	//fmt.Println(sum)
}
//帮我计算两个数的和，两个数的差
func cal2 (num1 int , num2 int) (int,int) {
	var sum int = 0
	sum += num1
	sum += num2


	var result int = num1 - num2
	return sum,result
}



func main(){
	//求两数的和： 10+20
	//调用函数：
	// sum := cal(10,20)
	// fmt.Println(sum)

	sum1,_ :=cal2(10,20)
	fmt.Println(sum1)
	
}