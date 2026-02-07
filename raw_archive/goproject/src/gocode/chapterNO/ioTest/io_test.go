package iotest

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestRead(t *testing.T) { //读取文件内容
	//io.ReadWriter()//将Read Writer做一个整合 以此类推
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4)
	var body []byte
	f.Seek(1, 0) //索引偏移一位进行读取
	for {
		n, err2 := f.Read(buf) //n为字节数
		if n == 0 || err2 == io.EOF {
			fmt.Println("读取完毕")
			break
		}
	}
	body = append(body, buf[:]...) //buf[:n]...
	fmt.Println(string(body))
}

func TestWrite(t *testing.T) { //向文件写入内容
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	f.Write([]byte("hello golang"))
	f.Close()

}

func TestCopy(t *testing.T) {
	//r := strings.NewReader("some io.Reader stream to be read\n")
	src, _ := os.Open("a.txt")
	src1, _ := os.OpenFile("test1.txt", os.O_CREATE|os.O_RDWR, 0664)
	if _, err := io.Copy(src, src1); err != nil { //将src1复制到src
		//os.Stdout:将内容输出到控制台
		log.Fatal(err) //直接退出程序
		//log.Fatal()函数完成;
		//1.打印输出err
		//2.退出程序
	}
}

func TestCopyBuffer(t *testing.T) { //与Copy的区别就是提供一个缓冲区(如果需要)
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func TestLimitReader(t *testing.T) { //从r读取,但限制数量为n个,每次读取更新n(可读取字节数量)
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4) //从r中读取4个

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func TestMultReader(t *testing.T) { //传入多个read,按逻辑依次读取read
	r1 := strings.NewReader("first reader")
	r2 := strings.NewReader("second reader")
	r3 := strings.NewReader("third reader")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func TestMultWrite(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2) //将其写入所有的写入器当中
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}

func TestPipe(t *testing.T) {
	r, w := io.Pipe() //创建一个管道r:read w:write
	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n") //将字符串写入W
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func TestReadAll(t *testing.T) { //直接读取所有
	r := strings.NewReader("Go is a general-purpose language designed with systems porgramming in mind.")
	b, err := io.ReadAll(r) //返回一个[]byte
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}

func TestReadAtLeast(t *testing.T) { //从r读取到buf,直至读取了min个字节,返回一个字节数
	r := strings.NewReader("some io.Reader strem to be read\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 3); err != nil { //至少读取3个
		log.Fatal(err)
	}
	fmt.Printf("buf: %s\n", buf)

	shortBuf := make([]byte, 3)                               //!
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil { //至少读取4个,但buf中只有3个,报错
		fmt.Println("error:", err)
	}

	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil { //至少读取64个,但源字符串没有64个，报错
		fmt.Println("error:", err) //unexpected EOF
	}
}

func TestReadFull(t *testing.T) {
	r := strings.NewReader("some io.Reader strem to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil { //只准确读取4个字节
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", buf)

	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil { //只准确读取64个字节,但字符串>64
		fmt.Println("error:", err)
	}
}

func TestNewSectionReader(t *testing.T) {
	r := strings.NewReader("some io.Reader strem to be read\n")
	s := io.NewSectionReader(r, 5, 17) //在r中从5读到17
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
}

func TestRead2(t *testing.T) {
	r := strings.NewReader("some io.Reader strem to be read\n")
	s := io.NewSectionReader(r, 5, 17) //io.Reader strem

	buf := make([]byte, 9)
	if _, err := s.Read(buf); err != nil { //从io.Reader strem中再读取9个
		log.Fatal(err)
	}
}

func TestReadAt(t *testing.T) {
	r := strings.NewReader("some io.Reader strem to be read\n")
	s := io.NewSectionReader(r, 5, 17)
	fmt.Println(s.Size()) //打印截取的字节大小
	buf := make([]byte, 6)
	if _, err := s.ReadAt(buf, 10); err != nil { //从索引偏移开始读，读6个
		//-------s.Seek(10,io.SeekStart)//偏移10个开始读,不需要buf
		//-------
		log.Fatal(err)
	}
}

func TestTeeReader(t *testing.T) { //向w写入从r读取到的内容
	var r io.Reader = strings.NewReader("some io.Reader strem to be read\n")

	r = io.TeeReader(r, os.Stdout)
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}
}

func TestWriteString(t *testing.T) {
	if _, err := io.WriteString(os.Stdout, "hello world"); err != nil { //将hello world 写入前面,w接受字节片
		log.Fatal(err)
	}
}
