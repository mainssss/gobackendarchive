package main

import (
	"fmt"
)

//存放数据
func numChan(intChan chan int) {
	for i := 1; i <= 2000; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Println("intChan协程运行完毕...")
}

//计算
func nubChan(intChan chan int, resChan chan int, boolChan chan bool) {
	//读取intChan
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		for i := 0; i <= v; i++ {
			v += i
			fmt.Println("v:", v)
		}
	}
	close(intChan)
	boolChan <- true
	close(boolChan)
	fmt.Println("协程运行完毕")
}

func main() {

	intChan := make(chan int, 2000) //存放数据
	resChan := make(chan int, 2000) //计算
	boolChan := make(chan bool, 8)
	go numChan(intChan)
	go func() {
		for i := 0; i <= 8; i++ {
			go nubChan(intChan, resChan, boolChan)
		}
	}()

	for {
		_, ok := <-boolChan
		if !ok {
			break
		}
	}

}
