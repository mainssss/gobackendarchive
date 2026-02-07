package main
import "fmt"
func main(){
	//定义一个map：
	b := make(map[int]string)
	//增加操作：
	b[20095452] = "张三"
	b[20095387] = "李四"
	//修改操作：
	b[20095452] = "王五"
	

	//删除操作：
	delete(b,20095387)
	delete(b,20089546)
	fmt.Println(b)
	
	//查找操作:
	value,flag := b[20092]
	fmt.Println(value)
	fmt.Println(flag)
}