package main
import (
	"fmt"
)

func main(){
	var age int = 18
	//&符号+变量 就可以获取这个变量内存的地址
	fmt.Println(&age)//0xc000010090

	//定义一个指针变量：
	//var代表声明一个变量
	//ptr 指针变量的名字
	//ptr对应的类型是：*int 是一个指针类型 （可以理解为 指向int类型的指针）
	//&age 就是一个地址，是ptr变量具体的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：",&ptr)

	//想获取ptr这个指针或者这个地址指向的数据
	fmt.Printf("ptr指向的数值为：%v",*ptr)  //ptr指向的数值为：18
}
//& 取内存地址
//* 根据地址取值