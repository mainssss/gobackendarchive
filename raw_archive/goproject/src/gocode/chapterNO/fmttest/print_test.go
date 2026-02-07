package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Print("我是控制台输出，不带换行\n", "第二个", 55555, "\n")
	fmt.Println("我是控制台输出，带换行", "第二个", 55555)
	fmt.Printf("我是格式化输出,%s占位符", "张三")
}

func TestFprint(t *testing.T) {
	fmt.Fprint(os.Stdout, "向标准输出\n")
	fmt.Fprintln(os.Stdout, "向标准输出")
	fmt.Fprintf(os.Stdout, "%s向标准输出\n", "张三")
}

func TestFprint1(t *testing.T) {
	s := os.FileMode(0777).String() //r：可读,w:可写,x:可执行
	//0777返回-rwxrwxrwx,0666返回-rw-rw-rw-,0644返回-rw-r--r--
	//-表示普通文件,d表示一个目录
	//2~4位:文件所有者权限rwx(可读/可写/可执行)
	//5~7位:文件所属用户组的权限rwx(可读/可写/可执行)
	//8~10位:其他人拥有的权限rwx(可读/可写/可执行)
	fmt.Println(s)
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%s向文件写入\n", "张三")
}

func TestSprint(t *testing.T) {
	host := "locahost"
	port := 6379
	add := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(add)
}

func TestErrorf(t *testing.T) {
	//errors.New("用户格式不正确:%d", host)  以下方法方便
	err := fmt.Errorf("用户名格式不正确:%s", "@#哈哈")
	if err != nil {
		panic(err)
	}
}
