package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个存在的文件，在原来的内容追加内容'ABC!ENGLISH'
	//1.打开文件 F:/abc.txt
	filePath := "F:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file err :%v\n", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入五句"hello Gardon"
	str := "ABC!ENGLISH!\r\n"
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
