package main

import "fmt"

//定义emp(雇员)
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//方法

//定义EmpLink,不带表头,即第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

//方法
//1.添加员工的方法,保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   //这是一个辅助节点
	var pre *Emp = nil //这是一个辅助指针 pre 在 cur 前面
	//如果当前的 EmpLink 就是一个空链表
	if cur == nil {
		this.Head = emp //完成
		return
	}

	//如果不是一个空链表,给emp找到对应的位置并插入
	//思路 让 cur 和 emp 比较，然后让pre始终保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置了
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时我们看一下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur

}

//显示当前链表信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 没有数据 \n", no)
		return
	}

	//遍历当前链表并显示数据
	cur := this.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行
}

//定义hashtable,本质是含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给hashtable编写insert 雇员的方法。
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数,确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//编写方法显示哈希表所有的雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对应链表的下标
}

func main() {

	id := 0
	key := ""
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("请选择菜单")
		fmt.Println("input表示添加雇员")
		fmt.Println("show表示显示雇员")
		fmt.Println("find表示查找雇员")
		fmt.Println("exit表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员名字")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
		case "exit":
		default:
			fmt.Println("输入有误")
		}
	}
}
