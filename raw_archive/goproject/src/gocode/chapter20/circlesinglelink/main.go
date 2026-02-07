package main

import (
	"fmt"
)

//定义猫的结构体
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成环状
		fmt.Println(newCatNode, "加入到环形链表")
		return
	}

	//定义一个临时变量,帮助我们找到环形的最后
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到列表中
	temp.next = newCatNode
	newCatNode.next = head
}

//输出环形的链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的具体情况:")
	temp := head
	if temp.next == nil {
		fmt.Println("空")
		return
	}

	for {
		fmt.Printf("猫的信息为:[id = %d,name = %s]->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	//将helper定位到环形链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	//如果只有一个结点
	if temp.next == head { //只有一个结点
		temp.next = nil
		return head
	}

	//如果有两个以上(包含2个)
	flag := true
	for {
		if temp.next == head { //如果到这，说明我已经比较到最后一个
			break
		}

		if temp.no == id {
			if temp == head { //说明删除的是头结点
				head = head.next
			}
			//找到(我们也可以在这里直接删除)
			helper.next = temp.next
			fmt.Printf("%d被删除\n", id)
			flag = false
			break
		}

		temp = temp.next     //移动【比较】
		helper = helper.next //移动【一旦找到要删除的结点】
	}

	//这里还要比较一次
	if flag == true { //如果flag为true,则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("%d被删除\n", id)
		} else {
			fmt.Printf("对不起,没有此id:%d\n", id)
		}
	}
	return head
}

func main() {
	//这里我们初始化一个环形链表
	head := &CatNode{}
	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head, 3)
	ListCircleLink(head)

}
