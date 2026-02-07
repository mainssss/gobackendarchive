package builtin

import (
	"fmt"
	"testing"
)

//此包不用引用,可以直接使用

func TestOne(t *testing.T) {
	s1 := []int{1, 2, 3}
	i := append(s1, 4) //将后面的silce添加到前面
	fmt.Printf("i: %v\n", i)

	s2 := []int{7, 8, 9}
	i2 := append(s1, s2...)
	fmt.Printf("i2: %v\n", i2)
}

func TestLen(t *testing.T) {
	s1 := "hello world"
	i := len(s1) //计算字节长度
	fmt.Printf("i: %v\n", i)

	s2 := []int{1, 2, 3}
	fmt.Printf("s2: %v\n", len(s2))
}
