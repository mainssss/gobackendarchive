package main
import "fmt"
func main(){
	//定义数组：
	var intarr[6]int = [6]int{1,4,7,2,5,8}
	//定义一个切片：
	var slice []int = intarr[1:4]
	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
	// fmt.Println(slice[2])
	// fmt.Println(slice[3])

	//对切片再次切片：
	slice2 := slice[1:2]
	fmt.Println(slice2)
	slice2[0] = 66
	fmt.Println(intarr)
	fmt.Println(slice)
}