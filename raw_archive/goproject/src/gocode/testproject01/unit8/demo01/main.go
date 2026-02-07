package main
import "fmt"

func main(){
	//定义数组:
	var intarr [6]int = [6]int{3,6,9,1,4,7}
	//切片构建在数组之上:
	//定义一个切片名字为slice，[]动态变化数组长度不写，int类型，intarr原数组
	//[1:3]切片-切出的一段片段-索引：从1开始,到3结束（不包含3）--[1,3)
	//var slice []int = intarr[1:3]
	slice := intarr[1:3]
	//输出数组：
	fmt.Println("intarr:",intarr)
	//输出切片：
	fmt.Println("slice:",slice)
	//切片元素个数:
	fmt.Println("slice的元素个数:",len(slice))
	//获取切片容量:容量可以动态变化
	fmt.Println("slice的容量为:",cap(slice))
	fmt.Printf("数组中下标为1位置的地址:%p",&intarr[1])
	fmt.Printf("切片中下标为0位置的地址:%p",&slice[0])
	slice[1] = 16
	fmt.Println("intarr:",intarr)
	fmt.Println("slice:",slice)
}