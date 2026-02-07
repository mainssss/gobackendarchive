package timetest

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //时
	minute := now.Minute() //分
	second := now.Second() //秒
	//02d输出的整数不足2位 用0补足
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTimeUnix(t *testing.T) {
	now := time.Now()
	fmt.Printf("cuuret time:%v\n", now)
	unixSec := now.Unix() //秒值
	fmt.Printf("unix sec value:%d\n", unixSec)
	milli := now.UnixMilli()
	fmt.Printf("unix milli value:%d\n", milli)
	//micro := now.UnixMicro()
	nowtime := time.UnixMilli(milli) //接收毫秒值算出当前时间
	fmt.Printf("cuuret time:%v\n", nowtime)

}

func TestParse(t *testing.T) {
	ti, err := time.Parse("2006-01-02 15:04:05", "2022-08-01 13:00:00") //解析时间
	if err != nil {
		panic(err)
	}
	fmt.Printf("t: %v\n", ti)

	loc, err := time.LoadLocation("Asia/Shanghai") //根据时区解析时间
	if err != nil {
		panic(err)
	}
	location, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-01 13:00:00", loc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("location: %s\n", location)
}

func TestFormat(t *testing.T) {
	now, err := time.Parse("2006-01-02 15:04:05", "2022-08-01 18:00:00") //解析时间
	if err != nil {
		panic(err)
	}
	//now := time.Now()
	format := now.Format("2006/01/02 15:04:05") //格式打印时间,24小时制
	format2 := now.Format("2006-01-02 ")
	format3 := now.Format("2006/01/02 03:04:05") //12小时化打印
	fmt.Printf("format: %v\n", format)
	fmt.Printf("format2: %v\n", format2)
	fmt.Printf("format2 12: %v\n", format2)
	fmt.Printf("format3: %v\n", format3)

}

func TestDuration(t *testing.T) {
	//five := 5 * time.Second //表示5秒
	fiveMinute := 5 * time.Minute //表示5分钟
	//add
	now := time.Now()
	fmt.Printf("now: %s\n", now)
	add := now.Add(fiveMinute) //表示当前时间加上上面的对应时间
	fmt.Printf("add: %s\n", add)
	//sub
	//sub := add.Sub(now) //当前时间减去对应时间
	// loc, err := time.LoadLocation("Asia/Shanghai") //根据时区解析时间
	// if err != nil {
	// 	panic(err)
	// }
	// parse, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-01 13:00:00", loc)
	// if err != nil {
	// 	panic(err)
	// }
	// sub := parse.Sub(now)
	// fmt.Printf("sub: %v\n", sub)
}
func TestComputer(t *testing.T) { //比较时间
	now := time.Now()
	now1 := time.Now()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(now.Equal(now1)) //比较时间
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-01 17:00:00", loc)
	t2, _ := time.Parse("2006-01-02 15:04:05", "2022-08-01 13:00:00")
	fmt.Println(t1, t2)
	fmt.Println(t1.Equal(t2))

	fmt.Printf("after:%t\n", t1.After(t2))   //代表t1的时间在t2之前
	fmt.Printf("before:%t\n", t1.Before(t2)) //代表t1的时间在t2之后
}

func TestTicker(t *testing.T) { //定时器
	// 	tick := time.Tick(5 * time.Second) //5秒钟之后执行一次
	// 	fmt.Println(time.Now())
	// 	for t2 := range tick {
	// 		fmt.Println(t2)
	// 	}
	// }
	// fmt.Println(time.Now())
	// time.Sleep(5 * time.Second)
	// time.AfterFunc(5*time.Second, func() {
	// 	fmt.Println("5秒钟之后才执行")
	// })
	for {

	}
}
