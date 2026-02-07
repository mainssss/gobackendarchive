package main
import "fmt"

func main(){
	//实现功能：如果口罩的库存小于30个，提示：库存不足，否则提示：库存充足
	//定义口罩的数量：
	var count int = 70
	if count < 30 { //如果这个条件表达式返回的是true的话，后面{}执行了
		fmt.Println("库存不足")
	} else {//count >= 30
		fmt.Println("库存充足")
	}

	//双分支一定会二选一走其中一个分支。
}
