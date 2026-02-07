package main

import (
	"fmt"
	"gocode/testroom/server/model"
	"net"
	"time"
)

/*
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的数据...")
	//conn.Read只有在conn没有被关闭的情况下才会阻塞
	//如果客户端关闭conn就不会阻塞了
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]换成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg  body err")
		return
	}

	//把pkg反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	return

}

func writePkg(conn net.Conn, data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail:", err)
		return
	}

	//2.发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail:", err)
		return
	}
	return
}
*/
//编写一个函数serverProcessLogin ,专门处理登录请求

/*
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.Data，并且反序列化成LoginMes:
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}

	//2.先声明一个resMes,并完成赋值
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//在声明一个LoginResMes
	var loginResMes message.LoginResMes

	//如果用户id为1425919357，密码为123456，认为是正确，否则不合法
	if loginmes.UserId == 1425919357 && loginmes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200

	} else {
		//不合法
		loginResMes.Code = 500 //500状态码表示该用户不存在
		loginResMes.Error = "该用户不存在,请注册在使用..."
	}

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	//6.发送data 我们将其封装到writePkg
	err = writePkg(conn, data)
	return

}
*/

//编写一个ServerProcessMes 函数
//功能:根据客户端发送的消息总类不同,决定调用哪个函数来处理
/*
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在,无法处理....")
	}
	return
}
*/

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭
	defer conn.Close()
	//这里调用总控t,创建一个总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程错误:", err)
		return
	}
}

//这里我们编写一个函数，完成对userDao的初始化任务
func initUserDao() {
	//这里的pool本身就是全局变量
	//这里需要注意初始化顺序问题
	//initPool,再initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	//当服务器启动时，我们就初始化我们的redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	//提示信息
	fmt.Println("(新)服务器在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	//一旦监听成功就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}

		//一旦连接成功则启动一个协程和客户端保持数据的交互(通讯)
		go process(conn)
	}
}
