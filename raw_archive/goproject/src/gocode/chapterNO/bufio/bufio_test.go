package bufio_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestBufio(t *testing.T) {
	s := strings.NewReader("ABCDEFG") //创建一个新reader
	str := strings.NewReader("123456")
	br := bufio.NewReader(s)    //NewReaderSize(rd,4096)
	b, _ := br.ReadString('\n') //按照换行来读取
	fmt.Println(b)
	br.Reset(str) //清除缓存区数据,下次读取使用str
	b, _ = br.ReadString('\n')
	fmt.Println(b)

	d := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br1 := bufio.NewReader(d) //创建一个缓存来读取
	p := make([]byte, 10)
	for {
		n, err := br1.Read(p)
		if n == 0 || err == io.EOF {
			break
		} else {
			fmt.Printf("string(p[0:n]):%v\n", string(p[0:n]))
		}
	}

}

func TestBufio2(t *testing.T) {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte() //读取一个字节
	//br.ReadRune()是读取中文
	//br.ReadLine()是按行读取
	fmt.Printf("c: %c\n", c)

	c, _ = br.ReadByte() //抛出上一个字节,读取一个字节
	fmt.Printf("c: %c\n", c)

	br.UnreadByte() //抛出字节
	c, _ = br.ReadByte()
	fmt.Printf("c: %c\n", c)
}

func TestBufio3(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	//bw.Available()返回剩余的字节
	//bw.Buffered()返回已经使用的字节
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c) //重置
	bw.WriteString("456")
	bw.Flush() //将缓存区的数据,写入下层的io.Writer接口(发牌)
	fmt.Println(b)
	fmt.Println(c)
}
