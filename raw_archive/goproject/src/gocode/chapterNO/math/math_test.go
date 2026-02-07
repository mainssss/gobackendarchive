package math_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestIsNaN(t *testing.T) { //判断是不是数值，如果是，返回false，反之
	fmt.Println(math.IsNaN(123321.321321))
}

func TestCeil(t *testing.T) { //向上取整
	fmt.Println(math.Ceil(1.13456))
}

func TestFloor(t *testing.T) { //向下取整
	fmt.Println(math.Floor(2.9999))
}

func TestTrunc(t *testing.T) { //取整数部分
	fmt.Println(math.Trunc(2.99999))
}

func TestAbs(t *testing.T) { //返回绝对值
	fmt.Println(math.Abs(2.99931232313214116655374))
}

func TestMax(t *testing.T) { //两数之间的最大值
	fmt.Println(math.Max(1000, 200))
}

func TestMin(t *testing.T) { //两数之间的最小值
	fmt.Println(math.Min(1000, 200))
}

func TestDim(t *testing.T) { //如果x-y<0则返回0,如果x-y>0则返回剩下的数
	fmt.Println(math.Dim(1000, 2000))
	fmt.Println(math.Dim(1000, 200))
}

func TestMod(t *testing.T) { //取余运算,可以理解为x-Trunc(x/y)*y,结果的正负号与x相同
	fmt.Println(math.Mod(123, 0))
	fmt.Println(math.Mod(123, 10))
}

func TestSqrt(t *testing.T) { //返回x的二次平方根
	fmt.Println(math.Sqrt(144))
}

func TestCbrt(t *testing.T) { //返回x的三次平方根
	fmt.Println(math.Cbrt(1728))
}

func TestHypot(t *testing.T) { //返回Sqrt(p*p+q*q),x的平方+y的平方根号
	fmt.Println(math.Hypot(12, 12))
}

func TestPow(t *testing.T) { //求幂，x的y次方
	fmt.Println(math.Pow(2, 3))
}

func TestSin(t *testing.T) { //求正弦
	fmt.Println(math.Sin(12))
}

func TestCos(t *testing.T) { //求余弦
	fmt.Println(math.Cos(12))
}

func TestTan(t *testing.T) { //求正切
	fmt.Println(math.Tan(12))
}

func TestLog(t *testing.T) { //求自然对数
	fmt.Println(math.Log(2))
}

func TestLog2(t *testing.T) { //求2为底数的对数
	fmt.Println(math.Log2(128))
}

func TestRand(t *testing.T) { //伪随机数
	rand.Seed(time.Now().Unix()) //设置种子，以当前时间的秒(也可以是纳秒)
	//可以在init()里面设置
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Intn(5)) //0~x内的随机数
}
