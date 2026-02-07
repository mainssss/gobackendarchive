package main

import (
	"errors"
	"fmt"
)

//使用数组模拟栈的使用
type Stack struct {
	MaxTop int //表示栈顶的最大可以存放的数
	Top    int //表示栈顶
	//因为栈底是固定的所以不给
	arr [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	//先判断栈是否已满

	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//遍历栈,注意需要从栈顶开始遍历
func (this *Stack) List() {
	//先判断这个栈是否为空
	if this.Top == -1 {
		fmt.Println("stack enpty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func (this *Stack) Pop() (val int, err error) {
	//先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("栈空!")
	}
	//先取值,再对this.top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func main() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //当栈顶为-1时表示栈为空

	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()
	val, _ := stack.Pop()
	fmt.Println("出栈val=", val)
	//显示
	stack.List()

	val, _ = stack.Pop()
	fmt.Println("出栈val=", val)
	//显示
	stack.List()
}
