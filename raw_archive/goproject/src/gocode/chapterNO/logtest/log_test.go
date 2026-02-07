package logtest

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) { //打印日志自带日期
	log.Print("this is a log")
	log.Printf("this is a log:%d", 100)
	name := "zhangsan"
	age := 20
	log.Println(name, "", age)
}

func TestLogPanic(t *testing.T) {
	defer fmt.Println("发生了panic错误!")
	log.Print("this is a log")
	log.Panic("this is a panic log") //抛出错误,但不会退出程序
	fmt.Println("运行结束")
}

func TestLogFatal(t *testing.T) {
	defer fmt.Println("发生了panic错误!")
	log.Print("this is a log")
	log.Fatal("this is a fatal log") //直接退出程序
	fmt.Println("运行结束")
}

func TestLogmore(t *testing.T) {
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //打印不同的东西
	log.Print("this is a log")
}

func TestLogmore2(t *testing.T) {
	s := log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.SetPrefix("[mylog]") //给log打印设置前缀
	s = log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.Print("this is a log") //[mylog]2022/10/23 21:19:04 this is a log
}

func TestLogflie(t *testing.T) {
	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f) //向文件写入
	log.Print("this is a log")
}
