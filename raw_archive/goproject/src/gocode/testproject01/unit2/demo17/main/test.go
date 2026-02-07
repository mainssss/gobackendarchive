package main //声明为main包，程序的入口包（所处于的目录下）

import (
	"fmt"
	"gocode/testproject01/unit2/demo17/test"
)

//程序的入口函数
func main(){
	//如果util.go中定义的是stuNo的话，那么在test.go中是访问不到的，会报错：.\test.go:10:14: cannot refer to unexported name test.stuNo
	fmt.Println(test.StuNo)
}