package testos

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCreate(t *testing.T) { //创建文件并清除
	//采用0666(任何人可读写,不可执行),创建一个名字为text.txt的文件，如果存在，会截断它(空文件)
	file, err := os.Create("test.txt")
	//pe := err.(*os.PathError)

	if err != nil {
		panic(err)
	}

	defer file.Close()
}

func TestMKdir(t *testing.T) { //创建单目录
	err := os.Mkdir("ms", os.ModePerm)
	//pe := err.(*os.PathError)

	if err != nil {
		panic(err)
	}
}

func TestMKdirAll(t *testing.T) { //创建单目录
	err := os.MkdirAll("ms/a/b", os.ModePerm)
	//pe := err.(*os.PathError)

	if err != nil {
		panic(err)
	}
}

func TestRemove(t *testing.T) { //删除一个空目录或文件
	err := os.Remove("ms/a/b")
	//pe := err.(*os.PathError)

	if err != nil {
		panic(err)
	}
}

func TestRemoveAll(t *testing.T) { //删除目录
	err := os.RemoveAll("ms")
	//pe := err.(*os.PathError)

	if err != nil {
		panic(err)
	}
}

func TestGetwd(t *testing.T) { //获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

func TestChDir(t *testing.T) { //更改当前工作目录
	err := os.Chdir("c:/goproject")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

func TestTempDir(t *testing.T) { //获取临时目录
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func TestRename(t *testing.T) { //重命名文件
	err := os.Rename("ms", "ms1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestFileStat(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat() //获取里面的文件信息，有名字，大小，修改时间等
	if err != nil {
		panic(err)
	}
	fmt.Println("flie size", fileInfo.Size())
	fmt.Printf("fileInfo: %v\n", fileInfo)
}

func TestRead(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var body []byte
	for {
		buf := make([]byte, 4)
		n, err := f.Read(buf) //读取文件
		if err == io.EOF {
			break
		}
		fmt.Printf("n:%d\n", n)
		body = append(body, buf[:n]...)
	}
	fmt.Printf("body: %s\n", body)
}

func TestReadAt(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 5)
	n, err := f.ReadAt(buf, 6) //从哪开始读
	fmt.Printf("buf:%s\n", buf[:n])
}

func TestReadDir(t *testing.T) {
	f, err := os.Open("testos")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dirs, err := f.ReadDir(-1) //读取目录(-1读取所有，2读取2个)
	for _, v := range dirs {
		fmt.Println(v.Name(), ":", v.IsDir())
	}

}

func TestFileSeek(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Seek(3, 0) //前一个是偏移位置,第二个是设置下一次读写的位置,0为开头,1为相对位置,2为文件结尾
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("读取内容:%s\n", buf[:n])
}

func TestFlieWrite(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("hello golang\n")) //两个都是写入
	f.WriteString("hello golang\n")
}

func TestFlieWriteAt(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteAt([]byte("insert content"), 5) //从5字节插入输入,将字节中的值进行替换
	fmt.Println(err)
}
