package main

import "fmt"

func main() {
	// //i表示层数，j表示每一层打印多少个星
	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// //打印半个金字塔
	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	//打印整个金字塔
	var scc int = 3
	for i := 1; i <= scc; i++ {
		for k := 1; k <= scc-1; k++ {
			fmt.Print(" ")
		}
		for j := 1; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
