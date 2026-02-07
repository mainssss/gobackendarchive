package main
import "fmt"

//函数功能：求和
//函数名字为getSum 参数为空
//getSum这个函数返回值为一个函数，并且这个函数的参数是int类型的参数，返回值也是int类型
func getSum() func (int) int {
	var sum int = 0
	return func (num int) int{
		sum = sum + num
		return sum
	}
}
//闭包：返回的匿名函数+匿名函数以外的变量num

func main(){
	f := getSum()
	fmt.Println(f(1))//1
	fmt.Println(f(2))//3
	fmt.Println(f(3))//6
	fmt.Println(f(4))//10

	fmt.Println("--------------------------------")
	fmt.Println(getSum01(0,1))//1
	fmt.Println(getSum01(1,2))//2
	fmt.Println(getSum01(3,3))//6
	fmt.Println(getSum01(6,3))//9
	
}

func getSum01(sum int,num int) int{
	sum = sum + num
	return sum
}

//不使用闭包的时候，我想保留的值，不可以反复使用
//闭包应用场景：闭包可以保留上次引用的某个值，我们传入一次就可以反复使用

