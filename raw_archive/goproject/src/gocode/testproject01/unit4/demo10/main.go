package main
import "fmt"

func main(){
	//双重循环:
	labele2:
	for i := 1; i <= 5;i++{
		for j := 2;j <= 4;j++{
			fmt.Printf("i:%v ,j:%v\n",i,j)
			if i == 2 && j == 2{
				break labele2 //结束指定标签对应的循环
			}
		}
	}
	fmt.Println("------OK")
	//如果标签声明没有使用，那么会报错
}
