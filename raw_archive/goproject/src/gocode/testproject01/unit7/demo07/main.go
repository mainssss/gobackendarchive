package main
import "fmt"

func main(){
	//定义一个二维数组:
	var arr [2][3]int16
	fmt.Println(arr)

	fmt.Printf("arr的地址是:%p",&arr)
	fmt.Printf("arr[0]的地址是:%p",&arr[0])
	fmt.Printf("arr[0][0]的地址是:%p",&arr[0][0])

	fmt.Printf("arr[1]的地址是:%p",&arr[1])
	fmt.Printf("arr[1][0]的地址是:%p",&arr[1][0])

	//赋值:
	arr[0][1] = 47
	arr[0][0] = 82
	arr[1][1] = 25
	fmt.Println(arr)

	//初始化操作:
	var arr1 [2][3]int = [2][3]int{{1,4,7},{2,5,8}}
	fmt.Println(arr1)
}