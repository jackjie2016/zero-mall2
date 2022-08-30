package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func main() {
	//time.Unix()
	logx.Error("测试的日志")

	//获取当前时间
	var t = time.Now()
	fmt.Println(&t)

	//格式化
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(&t)
	tm := time.Date(2020, time.April,
		11, 21, 34, 01, 0, time.UTC)
	fmt.Println(tm)
	//时间戳转换为本地时间
	fmt.Println("time.Unix:", time.Unix(1487780010, 1))
	//loc2, _ := time.LoadLocation("Local")
	todayZero, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Parse:", todayZero)

	//todayZero, _ = time.ParseInLocation("20060102150405", time.Now().Format("2006-01-02 15:04:05"), loc2)
	todayZero, _ = time.ParseInLocation("2006-01-02 15:04:05", "2015-01-01 00:00:00", time.Local)

	fmt.Println("ParseInLocation:", todayZero)
	//时区
	fmt.Println(t.Location())
	//返回时间点t对应的时间戳
	fmt.Println(t.Unix())
	//返回时间点t对应的年月日
	fmt.Println(t.Date())
	//输出小时
	fmt.Println(t.Hour())
	//输出分
	fmt.Println(t.Minute())
	//输出秒
	fmt.Println(t.Second())

	fmt.Println(t.Clock())

	toBeCharge := "2015-01-01 00:00:00" //待转化为时间戳的字符串 注意 这里的小时                                                                                和        分钟还要秒必须写 因为是跟着模版走                                                                            的 修改模板的话也可以不写

	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型

	sr := theTime.Unix()

	fmt.Println(theTime)
	fmt.Println(sr)

}

func time2int() {
	//t := &OldGood.CreatedAt

	//fmt.Println(OldGood.CreatedAt.Format("2006-01-02 15:04:05"))

	////转化所需模板
	//loc, _ := time.LoadLocation("Local")                                                                            //获取时区
	//theTime, _ := time.ParseInLocation(OldGood.CreatedAt.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05", loc) //使用模板在对应时区转化为time.time类型
	//
	//CreatedAt := time.Unix(int64(theTime.Unix()), 0)
	//
	//goods.CreatedAt = CreatedAt
}
