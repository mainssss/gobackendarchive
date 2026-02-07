package main
import "fmt"
func main(){
	var num1 int = 10
	fmt.Println(num1)
	var num2 int = (10 + 20) % 3 + 3 - 7   //=右侧的值运算完成后在赋值给左侧
	fmt.Println(num2)

	var num3 int = 10
	num3 += 20  //等价num3 = num3 + 20;
	fmt.Println(num3)

	//练习：交换两个数的值并输出结果
	var a int = 8
	var b int = 4
	fmt.Printf("a = %v,b = %v\n",a,b)
	//交换：
	//引入一个中间变量：
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("a = %v,b = %v",a,b)
}
