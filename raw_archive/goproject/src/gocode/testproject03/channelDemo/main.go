package main

import (
	"fmt"
)

func main() {
	//演示一下管道的使用
	//创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//intChan的类型
	fmt.Printf("intChan的值:%v intChan本身的地址:%p\n", intChan, &intChan)

	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 90 //注意：当我们给管道写入数据是不能超过make的内存

	//管道的长度和cap
	fmt.Printf("intChan len=%v,cap=%v\n", len(intChan), cap(intChan))

	//从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("intChan len=%v,cap=%v\n", len(intChan), cap(intChan))

	//在没有使用协程的情况下，如果管道数据取完还继续取，就会报错
	num3 := <-intChan
	num4 := <-intChan
	//num5 := <-intChan
	fmt.Printf("num3=%v  num4=%v\n", num3, num4)
}
