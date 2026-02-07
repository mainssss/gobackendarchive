package main
import "fmt"

func main(){
	//练习转义字符：
	// \n 换行
	fmt.Println("aaa\nbbb")
	// \b 退格
	fmt.Println("aaa\bbbb")
	// \光标回到本行开头，后续输入就会替换原有字符
	fmt.Println("aaaaa\rbbb")
	// \t 制表符
	fmt.Println("aaaaaaaaa")
	fmt.Println("aaaaaa\tbbbbb")
	fmt.Println("aaaaaaaaa\tbbbbb")

	//\"
	fmt.Println("\"Golang\"")
}