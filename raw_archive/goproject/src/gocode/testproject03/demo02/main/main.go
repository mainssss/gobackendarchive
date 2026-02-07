package main

import (
	"fmt"
	"os"
)

func main() {
	//打开一个文件
	//概念说明:file的叫法
	//1.file叫file对象
	//2.file叫file指针
	//3.file叫file文件句柄
	file, err := os.Open("E:/test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}

	//输出文件,看出file就是指针
	fmt.Printf("file = %v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件错误:", err)
	}
}
