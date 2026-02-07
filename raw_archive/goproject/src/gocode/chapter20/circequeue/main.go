package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int // 4
	array   [4]int
	head    int //指向队首 0
	tail    int //指向队尾 0
}

//入队列 AddQueue(push) GetQueue(pop)

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	//分析出尾部this.tail 在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val //把值给尾部
	this.tail++
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpt() {
		return 0, errors.New("queue empty")
	}
	//取出,head 是指向队首，并且包含队首元素
	val = this.array[this.head]
	this.head++
	return
}

//显示队列

//判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断环形队列是否为空
func (this *CircleQueue) IsEmpt() bool {
	return this.tail == this.head
}

//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	//这是一个关键算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数")
			fmt.Scanln(&val)
			/*
				err := queue.AddQueue(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {

					fmt.Println("加入队列ok")
				}*/
		case "get":
		/*
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数:", val)
			}
		*/
		case "show":
			//queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}

	}
}
