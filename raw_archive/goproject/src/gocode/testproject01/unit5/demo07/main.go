package main
import "fmt"

//参数类型为指针：
func test(num *int){
	//对地址对应的变量进行改变数值：
	*num = 30
}

func main(){
	var num int = 10
	fmt.Println(&num)
	test(&num) //调用test函数的时候，传入的是num的地址
	fmt.Println(num)
}	

