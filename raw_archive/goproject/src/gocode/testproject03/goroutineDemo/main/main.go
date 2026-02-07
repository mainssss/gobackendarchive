package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1~200各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来，要求使用goroutine

// 思路
// 1.编写一个函数来计算各个数的阶乘，并放入到map中
// 2.我们启动多个协程，统一将结果放入map中
// 3.map应该做成全局的

var (
	MyMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是全局的互斥锁
	//sync是包，全称是synchornzed 同步的意思
	//Mutex是互斥的意思
	lock sync.Mutex
)

//test函数就是计算n，将结果放入map中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//我们将res结果放入到myMap中
	//加锁
	lock.Lock()
	MyMap[n] = res //concurrent map writes
	//解锁
	lock.Unlock()
}

func main() {

	//我们这里开启多个协程完成任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠10秒钟[等多久？]
	time.Sleep(time.Second * 5)

	//这里我们遍历这个结果
	lock.Lock()
	for i, v := range MyMap {
		fmt.Printf("map[%d]%d\n", i, v)
	}
	lock.Unlock()
}
