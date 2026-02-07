package main
import (
	"fmt"
	"errors"
)

func main(){
	err := test()
	if err != nil{
		fmt.Println("自定义错误：",err)
		panic(err)
	}
	fmt.Println("除法操作执行成功-----")
	fmt.Println("正常执行下面的逻辑------")
}

func test()(err error){
	num1 := 10
	num2 := 0
	if num2 == 0{
		//抛出自定义错误
		return errors.New("除数不能为0")
	}else{//如果除数不为0，那么正常执行就可以了
		result := num1 / num2
		fmt.Println(result)
		//如果没有错误的话，返回零值
		return nil
	}
	
}