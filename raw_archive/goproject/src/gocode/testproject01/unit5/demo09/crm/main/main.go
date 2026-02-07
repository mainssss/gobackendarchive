package main //1·packeage进行包的声明，建议：包的声明这个包与所在的文件夹同名
//2.main包是程序的入口包，一般main函数会放在这个包下
import (
	"fmt"
	"gocode/testproject01/unit5/demo09/crm/dbutils"
	test "gocode/testproject01/unit5/demo09/crm/calutils"
)
func main(){
	fmt.Println("你好,这是main函数的执行")
	bbb.GetConn()//4.在函数调用的时候前面需要定位到所在的包

	test.Add()
}