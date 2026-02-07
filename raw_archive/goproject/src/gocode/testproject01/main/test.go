package main //声明文件所在包（每个go文件必须有归属的包）
import "fmt" //引入程序中需要用的包，为了使用包下的函数 比如：Println
func main() { //main程序主函数 程序主入口
	fmt.Println("Hello Golang!") //在控制台打印输出一句话，双引号中的内容会原样输出
	fmt.Println("Hello Golang!")
	fmt.Println("Hello Golang!")
	fmt.Println("Hello Golang!")
	var age = 18 + 10
	fmt.Println(age)

}
