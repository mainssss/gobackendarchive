package main

import (
	"fmt"
	"net" //在socket开发中，net包是非常重要的
)

func process(conn net.Conn) {
	//这里循环接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有wrtie，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取，放入buf中
		if err != nil {

			fmt.Println("客户端退出 err=", err)
			return
		}
		//3.显示客户端发送的内容到终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	//1.tcp表示使用网络协议是tcp
	//2. 0.0.0.0：8888表示在本地监听8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Listen err=", err)
		return
	}

	defer listen.Close() //延时关闭listen

	//循环等待客户端连接
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept() //Accept()有客户端连接才进行下一步
		if err != nil {
			fmt.Println("Accept() err=", err)
			return
		} else {
			fmt.Printf("Accept() suc con=%v 客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备起一个协程为客户端服务
		go process(conn)
	}

	//fmt.Printf("listen suc=%v", listen)
}
