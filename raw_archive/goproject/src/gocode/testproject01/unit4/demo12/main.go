package main
import "fmt"

func main(){
	//双重循环:
	labele:
	for i := 1; i <= 5;i++{
		for j := 2;j <= 4;j++{
			if i == 2 && j == 2{
				continue  labele// 结束离它近的循环,继续下一次循环
			}
			fmt.Printf("i:%v ,j:%v\n",i,j)
		}
	}
	fmt.Println("------OK")
	//如果标签声明没有使用，那么会报错
}
