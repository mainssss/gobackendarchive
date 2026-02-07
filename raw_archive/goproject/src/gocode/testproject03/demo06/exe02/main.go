package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//创建一个新文件，写入内容5句"hello Gardon"
	//1.打开文件 F:/abc.txt
	filePath := "F:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file err :%v\n", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入五句"hello Gardon"
	str := "hello golang\r\n"
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriteString方法时，
	//其实内容是先写入到缓存的,所以需要调用Flush方法，将缓存的数据真正写入到文件中
	//否则文件中会没有数据
	writer.Flush()

}
