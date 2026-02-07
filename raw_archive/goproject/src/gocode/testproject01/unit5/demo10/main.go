package main
import (
	"fmt"
	"gocode/testproject01/unit5/demo10/testutils"
)
var num int = test()

func test() int{
	fmt.Println("test函数被执行")
	return 10
}

func init(){
	fmt.Println("main.init函数被执行!")
}

func main(){
	fmt.Println("main函数被执行!")
	fmt.Println("Age=",testutils.Age,"Sex=",testutils.Sex,"Name=",testutils.Name)
}