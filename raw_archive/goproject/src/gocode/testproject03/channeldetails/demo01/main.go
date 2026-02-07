package main

import (
	"fmt"
)

func main() {
	//管道可以声明为只读或者只写

	//1.默认情况下管道可以声明为只读/只写
	//var chan1 chan int //可读可写

	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	fmt.Println("chan2=", chan2)

	//3.声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	fmt.Println("num2=", num2)
}
