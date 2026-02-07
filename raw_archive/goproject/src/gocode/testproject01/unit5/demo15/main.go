package main
import (
	"fmt"
	"strings"
)
func main(){
	//字符串的替换：
	str1 := strings.Replace("goandjavagogo","go","golang",-1)
	str2 := strings.Replace("goandjavagogo","go","golang",2)
	fmt.Println(str1)
	fmt.Println(str2)

	//按照指定的某个字符，为分割标识，将一个字符串拆封成字符串数组：
	arr := strings.Split("go-python-java","-")
	fmt.Println(arr)

	//将字符串进行大小写转换：
	fmt.Println(strings.ToLower("Go"))//go
	fmt.Println(strings.ToUpper("go"))//Go

	//将字符串左右两边空格去掉：
	fmt.Println(strings.TrimSpace("     go and java     "))

	//将字符串左右两边的字符去掉：
	fmt.Println(strings.Trim("~golang~","~"))

	//将左边指定字符串去掉：
	fmt.Println(strings.TrimLeft("~golang~","~"))

	//将右边指定字符串去掉：
	fmt.Println(strings.TrimRight("~golang~","~"))

	//判断字符串是否以指定的字符串开头：
	fmt.Println(strings.HasPrefix("http://java.sun.com/jsp/jst1/fmt","http"))

	//判断字符串是否以指定字符串结束：
	fmt.Println(strings.HasSuffix("demo.png",".jpg"))
}



