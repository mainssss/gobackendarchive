package main
import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	//统计字符串的长度，按字节进行统计
	str := "Golang你好"//在go语言中汉字是utf-8字符集，一个汉字3个字节
	fmt.Println(len(str))	//12

	//对字符串进行遍历：
	//方式1:利用键值循环： for-range
	for i, value := range str {
		fmt.Printf("索引为:%d,具体的值为%c \n",i,value)
	}

	//方式2:利用r:=[]rune(str)
	r := []rune(str)
	for i := 0; i <len(r);i++{
		fmt.Printf("%c\n",r[i])
	}

	//字符串转整数：
	num1,_ := strconv.Atoi("666")
	fmt.Println(num1)

	//整数转字符串：
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	//统计一个字符串中有几个子串：
	count := strings.Count("golangandjavaga","ga")
	fmt.Println(count)

	//不分大小写的字符串比较：
	flag := strings.EqualFold("hello","Hello")
	fmt.Println(flag)

	//区分大小写进行字符串比较：
	fmt.Println("hello"=="Hello")

	//返回子串在字符串第一次出现的索引值，如果没有返回-1
	fmt.Println(strings.Index("golangandjavaga","ga0"))
}



