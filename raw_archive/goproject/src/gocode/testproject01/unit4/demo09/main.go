package main
import "fmt"

func main(){
	//功能：求1~100的和，当和第一次超过300的时候，停止程序
	var sum int = 0
	for i := 1 ;i <= 100 ; i++{
		sum += i
		fmt.Println(sum)
		if sum >= 300{
			//停止正在执行的这个循环：
			break
		}
	}
	fmt.Println("----------END")
}
