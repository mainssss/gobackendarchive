package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {

	//将store序列化
	data, err := json.Marshal(&this)
	if err != nil {
		fmt.Printf("序列化错误 err= %v\n", err)
		return false
	}

	//保存文件
	filePath := "e:/monst.txt"
	err = ioutil.WriteFile(filePath, data, 066)
	if err != nil {
		fmt.Printf("witer file err= %v\n", err)
		return false
	}
	return true
}

//给monster定义一个方法ReStore，可以将一个序列化的monster从文件中读取
//反序列化为monster对象，检查反序列化，名字正确

func (this *Monster) ReStore() bool {
	//先读取文件中字符串
	filePath := "e:/monst.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file err =%v", err)
		return false
	}
	//用读取到的data []byte 进行反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("unmarshal err=", err)
		return false
	}
	return true
}
