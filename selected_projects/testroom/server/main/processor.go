package main

import (
	"fmt"
	"gocode/testroom/common/message"
	process2 "gocode/testroom/server/process"
	"gocode/testroom/server/utils"
	"io"
	"net"
)

//现创建一个Processer 结构体
type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		//创建一个UserProcessor 实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在,无法处理....")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环读取客户端发送的信息
	for {

		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回mes,err
		//创建一个Transfer 实例，完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器端也关闭")
				return err
			}
			fmt.Println("readPkg err", err)
			return err
		}
		//fmt.Println("mes:", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
