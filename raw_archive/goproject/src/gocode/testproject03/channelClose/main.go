package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//这时候不能在写入数据到channel
	//intChan <- 300
	fmt.Println("ooo")
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	//遍历管道不能使用普通的for循环
	//如果遍历时channel没有关闭，代码则会报错deadlock
	//关闭管道
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
