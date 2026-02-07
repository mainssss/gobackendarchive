package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "E:/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("错误:%v", err)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v", content) //[]byte
	fmt.Printf("%s", content)
	//因为我们没有显示的Open文件，因此也不需要显式的Close文件
	//因为，文件的Open和Close被封装到ReadFile函数内部
}
