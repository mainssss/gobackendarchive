package main
import "fmt"
func main(){
	//new分配内存，new函数的实参是一个类型而不是具体数值，new函数返回值是对应类型的指针 num : *int
	num := new(int)
	fmt.Printf("num对应的类型是: %T,num的值是:%v,num的地址:%v,num指针指向的值是:%v",num,num,
	&num,*num)
}