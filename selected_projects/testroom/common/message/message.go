package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

// 定义两个消息，后面需要在增加
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //状态码，500表示该用户还没注册，200表示登录成功
	Error string `json:"error"` //返回的错误信息，如果正确就=nil
}

type RegisterMes struct {
	User User `json:"user"` //类型就是user结构体.
}
type RegisterResMes struct {
	Code  int    `json:"code"`  //状态码，400表示该用户已经存在，200表示登录成功
	Error string `json:"error"` //返回的错误信息，如果正确就=nil
}
