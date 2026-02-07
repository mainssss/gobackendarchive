package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //如果文件或者目录存在
		fmt.Println("文件存在")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return false, nil
	}
	return false, err
}
func main() {
	PathExists("F:/abc.txt")
}
