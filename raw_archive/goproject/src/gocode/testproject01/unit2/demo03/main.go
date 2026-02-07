package main
//import "fmt"
//import "unsafe"
import(
	"fmt"
	"unsafe"
)

func main(){
	//定义一个整数类型：
	var num1 int8 = 120
	fmt.Println(num1)

	var num2 uint8 = 200
	fmt.Println(num2)

	var num3 = 28
	//实际上printf函数的作用是：格式化，把num3的类型填充到%t的位置上
	fmt.Printf("num3的类型是： %T", num3)//num3的类型是： int
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num3))

	//表示学生年龄：
	var age byte = 18
	fmt.Println(age)
}