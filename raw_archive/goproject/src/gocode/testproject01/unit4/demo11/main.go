package main
import "fmt"

func main(){
	//功能：输出1~100中被6整除的数：
	//方式1
	// for i := 1 ; i<=100 ; i++ {
	// 	if i % 6 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//方式2：
	for i := 1 ; i<=100 ; i++ {
			if i % 6 != 0 {
				continue //结束本次循环，继续下一次循环
			}
			fmt.Println(i)
		}
}
