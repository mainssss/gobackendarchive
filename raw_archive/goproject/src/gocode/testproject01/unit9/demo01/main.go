package main
import "fmt"
func main(){
	//方式1：
	//定义一个map的变量：
	var a map[int]string
	//只声明map，内存没有分配空间：
	//必须通过make函数进行初始化，才会分配空间：
	a = make(map[int]string,10)//map可以存放10个键值对
	//将键值对存入map中:
	a[20095452] = "张三"
	a[20095387] = "李四"
	//输出集合
	fmt.Println(a)

	//方式2：
	b := make(map[int]string)
	b[20095452] = "张三"
	b[20095387] = "李四"
	fmt.Println(b)

	//方式3：
	c := map[int]string{
		20095452:"张三",
		20098765:"李四",
	}
	c[20095387] = "王五"
	fmt.Println(c)
}