package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}

	//关闭
	close(intChan)
}

//向intChan取出数据,并判断是否为素数,如果是，就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环
	var flag bool
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { //intChan 取不到
			break
		}
		flag = true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数放入primeChan
			primeChan <- num
		}

	}
	fmt.Println("primeNum 因为取不到数据退出...")
	//这里还不能关闭primeChan
	//向exitChan写入true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个

	start := time.Now().Unix()

	//开启一个协程，向intChan放入1~8000个数
	go putNum(intChan)
	//开启4个协程，向intChan取出数据,并判断是否为素数,如果是，就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程要进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)

		//当我们从这个管道取出4个结果，我们就可以放心的关闭primeNum
		close(primeChan)
	}()

	//遍历我们的primeNum，把结果取出
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数:%d\n", res)
	}

	fmt.Println("main主线程退出...")
}
