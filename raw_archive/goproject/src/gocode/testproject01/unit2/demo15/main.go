package main
import (
	"fmt"
)

func main(){
	var num int = 10
	fmt.Println(num)

	var ptr *int = &num
	*ptr = 20
	fmt.Println(num)

}
