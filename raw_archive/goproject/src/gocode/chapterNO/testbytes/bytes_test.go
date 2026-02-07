package testbytes

import (
	"bytes"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var b = []byte("seafood")
	//upper := bytes.ToUpper(b) //转为大写
	upper := bytes.ToLower(b) //转为小写
	fmt.Println(string(b), string(upper))
}

func TestBytes(t *testing.T) {
	var b = []byte("ABC")
	var a = []byte("AbC")
	ret := bytes.Compare(a, b) //比较ab是否一样,相等返回0，a<b返回-1,a>b返回1
	fmt.Printf("compare result:a=%v b=%v,%d\n", a, b, ret)
	boo := bytes.Equal(a, b)
	fmt.Printf("equal result:a=%v b=%v,%t\n", a, b, boo) //判断是否相等
	boo = bytes.EqualFold(a, b)                          //忽略大小写比较
	fmt.Printf("EqualFold result:a=%v b=%v,%t\n", a, b, boo)
}

func TestBytes1(t *testing.T) {
	bs := [][]byte{
		[]byte("Hello world !"),
		[]byte("Hello 世界！"),
		[]byte("Hello golang."),
	}

	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("! ！ ."), r) //判断r字符是否包含"!！."
	}

	for _, b := range bs {
		fmt.Printf("去掉两边:%q\n", bytes.TrimFunc(b, f)) //去掉两边满足函数的字符
		fmt.Printf("去掉两边:%q\n", bytes.Trim(b, "!"))   //去掉两边满足函数的字符
	}

	for _, b := range bs {
		fmt.Printf("去掉前缀:%q\n", bytes.TrimPrefix(b, []byte("Hello"))) //去前缀
	}
}

func TestBytes2(t *testing.T) {
	b := []byte("   Hello   World !   ")
	fmt.Printf("b:%q\n", b)
	fmt.Printf("%q\n", bytes.Split(b, []byte{' '})) //使用空格字符进行分隔["" "" "" "Hello" "" "" "World" "!" "" "" ""]
	fmt.Printf("%q\n", bytes.Fields(b))             //以连续空格进行分隔["Hello" "World" "!"]
	f := func(r rune) bool {

		return bytes.ContainsRune([]byte(" !"), r) //满足函数中的内容进行分隔,结果不包含分隔符
	}
	fmt.Printf("%q\n", bytes.FieldsFunc(b, f)) //["Hello" "World"]

	bs := [][]byte{
		[]byte("Hello world !"),
		[]byte("Hello 世界！"),
		[]byte("Hello golang."),
	}
	join := bytes.Join(bs, []byte("--------")) //以内容为连接符号,将bs连成一个字节串
	fmt.Printf("%s\n", join)
}

func TestBytes3(t *testing.T) {
	b := []byte("hello golang")
	sublice1 := []byte("hello")
	sublice2 := []byte("Hello")
	fmt.Println(bytes.Contains(b, sublice1)) //b中是否包含sublitce1
	fmt.Println(bytes.Contains(b, sublice2)) //b中是否包含sublitce2

	s := []byte("hellooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1))              //s中包含几个sep1
	fmt.Println(bytes.Count(s, sep2))              //s中包含几个sep2
	fmt.Println(bytes.Count(s, sep3))              //s中包含几个sep3
	fmt.Println(bytes.HasPrefix(s, []byte("hel"))) //是否包含前缀
	fmt.Println(bytes.Index(s, []byte("o")))       //索引"o"第一次出现的位置
}

func TestBytes4(t *testing.T) {
	s := []byte("hello world")
	old := []byte("o")
	news := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  //将s中的old替换成news，替换0次
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  //将s中的old替换成news，替换1次
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  //将s中的old替换成news，替换2次
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //将s中的old全部替换成news

	s1 := []byte("你好世界")
	r := bytes.Runes(s1)
	fmt.Println("转换前字符串的长度:", len(s1)) //12(一个中文占4个字节)
	fmt.Println("转换后字符串的长度:", len(r))  //4
	//主要用于计算字节(按那种方式计算)
}
