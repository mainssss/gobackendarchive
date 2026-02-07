package main

import (
	"fmt"
)

func main() {
	n := 4
	test(n)
}

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println(n)
	}

}
