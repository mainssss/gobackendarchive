package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/testroom/cilent/utils"
	"gocode/testroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器段
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接失败:", err)
		return
	}

	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//将 resgisterMes 序列化
	data, err := json.Marshal(registerMes) //data是一个[]byte
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	data, err = json.Marshal(mes) //data是一个[]byte
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册信息发送错误 err:", err)
	}
	mes, err = tf.ReadPkg() //mes就是registerResMes

	if err != nil {
		fmt.Println("readPkg(conn) err:", err)
		return
	}
	//将mes的Data反序列化成registerResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,请重新登录")

		os.Exit(0)

	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return

}

//关联一个用户登录的方法
//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//下一步开始定协议
	//fmt.Printf("userId=%d\nuserPwd=%s", userId, userPwd)
	//return nil

	//1.连接到服务器段
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接失败:", err)
		return
	}

	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes) //data是一个[]byte
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes) //data是一个[]byte
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	//7.data是要发送的数据
	//7.(1)先把data长度发送给服务器
	//先获取到data长度，然后将长度转成一个可以表示长度的[]byte切片
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
	fmt.Printf("客户端,发送消息长度=%d 内容=%s", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail:", err)
		return
	}

	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20...")
	//这里还需要处理服务器端返回的消息.
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() //mes就是
	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}
	//将mes的Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")
		//这里我们起一个协程
		//该协程保持和服务器的通讯，如果服务器有数据推送给客户端
		//则可以接收并显示在客户端的终端

		go serverProcessMes(conn)
		//1.显示我们登录成功后的菜单[循环]..
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
