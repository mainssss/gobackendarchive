package main
import "fmt"

//自定义函数：功能：交换两个数
func exchangeNum (num1 int , num2 int){
	var t int
	t = num1
	num1 = num2 
	num2 = t
}

func main(){
	//调用函数：交换10和20
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数： num1 = %v,num2 = %v \n",num1,num2)
	exchangeNum(num1,num2)
	fmt.Printf("交换后的两个数： num1 = %v,num2 = %v \n",num1,num2)
}	
//运行结果：
//交换前的两个数： num1 = 10,num2 = 20
//交换后的两个数： num1 = 10,num2 = 20

