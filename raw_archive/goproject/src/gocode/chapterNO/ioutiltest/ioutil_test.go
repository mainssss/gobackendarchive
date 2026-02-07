package ioutil_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestNopClose(t *testing.T) { //把打开的文件与close进行封装
	f, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}

	ioutil.NopCloser(f)
}

func TestReadAll(t *testing.T) { //从r读取数据直到遇到EOF或者err，返回读取到的数据和可能的错误,成功调用返回的err为nil
	f, _ := os.Open("a.txt")
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b: %v\n", string(b))
}

func TestReadDir(t *testing.T) { //读取目录内的所有信息
	f, _ := ioutil.ReadDir(".")
	for _, v := range f {
		fmt.Printf("v: %v\n", v.Name())
	}
}

func TestReadFile(t *testing.T) { //打开并读取
	b, _ := ioutil.ReadFile("a.txt")
	fmt.Printf("b: %v\n", string(b))
}

func TestWriteFile(t *testing.T) { //打开并使用切片(缓存)写入文件(清空写入),再按照权限操作文件
	ioutil.WriteFile("a.txt", []byte("hello golang"), 0644)
}

func TestTempDir(t *testing.T) { //在dir目录中、创建一个新的使用prfix作为前缀的文件夹,并返回文件夹的路径
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dir: %v\n", dir)

	tempfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tempfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}
