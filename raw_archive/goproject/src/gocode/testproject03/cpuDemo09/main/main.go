package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum =", cpuNum)

	//可以自己设置使用多少个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("OK")
}
