package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将F:/abc.txt文件内容导入到E:/kkk.txt
	//1.首先将F:/abc.txt内容读取到内存
	//2 将读取到的内容写入E:/kkk.txt
	file1Path := "F:/abc.txt"
	file2Path := "E:/kkk.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		//说明读取文件有错误
		fmt.Println("read file err =", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file error = %v", err)
	}
}
