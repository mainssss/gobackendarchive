package testutils
import "fmt"

var Age int
var Sex string
var Name string

//定义一个init函数对变量进行初始化赋值：
func init(){
	fmt.Println("testutils中的init函数被执行了")
	Age = 19
	Sex ="女"
	Name = "lili"
}