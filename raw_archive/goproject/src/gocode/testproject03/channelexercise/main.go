package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {

	//定义一个可以存放任何数据类型的管道
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小猫", 4}
	allChan <- cat

	//我们希望获取到管道中的第三个元素则先推出前两个
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的cat是什么
	fmt.Printf("newCat=%T newCat=%v", newCat, newCat)
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
