package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，接收两个文件路径 srcFileName dstFeilName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err =", err)
	}
	defer srcFile.Close()
	//通过strfile句柄，获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFeilName
	destFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err :%v\n", err)
		return
	}
	//通过dstFeil获取到written
	writer := bufio.NewWriter(destFile)
	defer destFile.Close()
	return io.Copy(writer, reader)
}
func main() {
	//将d:/lop.jpg 文件拷贝到 e:/abc.jpg
	//调用CopyFile 完成文件拷贝
	srcFile := "d:/lop.jpg"
	dstFile := "e:/abc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝错误:", err)
	}

}
