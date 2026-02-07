package main

import (
	"fmt"
)

//小孩结构体
type Boy struct {
	no   int
	next *Boy
}

//编写一个函数，构成单向环形链表
//num:表示小孩的个数
//返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {

	first := &Boy{}
	curBoy := &Boy{}
	//判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环构建环形链表
	for i := 1; i <= num; i++ {

		boy := &Boy{
			no: i,
		}
		//构成循环链表需要一个辅助指针
		//1.因为第一个小孩比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first //构成环形链表
		}
	}
	return first
}

//显示单向的环形链表
func showBoy(first *Boy) {
	//先创建一个指针帮助遍历
	curBoy := first
	//处理一下如果环形链表为空的情况
	if first.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		//说明至少有一个小孩
		fmt.Printf("小孩编号=%d ->", curBoy.no)
		//退出条件
		if curBoy.next == first {
			break
		}
		//下移
		curBoy = curBoy.next
	}
}

func main() {
	first := AddBoy(5)
	showBoy(first)
}
