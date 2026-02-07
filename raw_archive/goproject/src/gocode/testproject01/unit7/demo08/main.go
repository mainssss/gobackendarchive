package main
import "fmt"

func main(){
	//定义一个二维数组:
	var arr [3][3]int = [3][3]int{{1,4,7},{2,5,8},{3,6,9}}
	fmt.Println(arr)
	fmt.Println("------------------------------")

	//方式1：普通for循环
	for i := 0 ; i < len(arr) ; i++{
		for j := 0 ; j < len(arr[i]) ; j++{
			fmt.Print(arr[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println("------------------------------")
	
	//方式2：for   range循环
	for key,value := range arr {
		for k,v := range value {
			fmt.Printf("arr[%v][%v]=%v\t",key,k,v)
		}
		fmt.Println()
	} 
}