package cal

import (
	//"fmt"
	//"log"
	"fmt"
	"testing" //引入go 的testing框架包
)

//编写一个测试用例,去测试addUpper()是否正确
func TestAddUpper(t *testing.T) {

	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("addUpper(10)执行错误,期望值=%v 实际值=%v\n",55,res)
		t.Fatalf("addUpper(10)执行错误,期望值=%v 实际值=%v\n", 55, res)
	}

	//如果正确，输出日志
	t.Logf("addUpper(10)执行正确")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用")
}
