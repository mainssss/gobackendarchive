package main
import (
	"fmt"
	"strconv"
)

func main(){
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1),10) //参数：第一个参数必须转为int64类型，第二个参数字面值的进制形式为十进制
	fmt. Printf("s1对应的类型是:%T,s1 = %q \n",s1, s1)

	var n2 float64 = 4.29
	var s2 string = strconv.FormatFloat(n2,'f',9,64)
	//第二个参数：'f'(-ddd.dddd) 第三个参数：9 保留小数点后9位 第四个参数指的是：表示当前小数是float64类型
	fmt. Printf("s2对应的类型是:%T,s2 = %q \n",s2, s2)

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt. Printf("s3对应的类型是:%T,s3 = %q \n",s3, s3)
}