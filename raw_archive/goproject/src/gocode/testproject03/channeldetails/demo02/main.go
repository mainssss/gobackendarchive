package main

import (
	"fmt"
	"time"
)

func main() {

	//使用select可以解决从管道取数据的阻塞问题

	//1.定义一个管道 10个数据int类型
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string类型
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭，会阻塞，而导致死锁

	//问题：在实际开发中，可能我们不好确定什么时候关闭该挂到.
	//可以使用select方式可以解决
	for {
		//注意：这里如果管道最终一直没有关闭，不会一直阻塞而死锁
		//会向下自动到下一个case匹配
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据:%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取到的数据:%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都取不到了,不玩了,程序员可以加入自己的逻辑\n")
			time.Sleep(time.Second)
			return
		}

	}
}
