package tools

import (
	"fmt"
	"strings"
	"time"
)

// 获取时间尝试方式
func timeTool1() {
	now := time.Now() // 获取当前时间，返回 Time 类型
	fmt.Println("当前时间：", now)

	// 基于 Time 获取年月日时分秒
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 获取时间戳
	fmt.Println("当前时间戳：", now.Unix())
}

// 时区问题
func timeTool2() {
	// 自定义一个时区
	hangzhou := time.FixedZone("Hangzhou", int((8 * time.Hour).Seconds()))
	// 获取内置时区
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("load Asia/Tokyo location failed", err)
		return
	}
	// 获取时间
	timeHangzhou := time.Date(2022, 4, 6, 21, 37, 48, 0, hangzhou)
	timeTokyo := time.Date(2022, 4, 6, 21, 37, 48, 0, tokyo)
	fmt.Println(timeHangzhou)
	fmt.Println(timeTokyo)
	// 看似时间一样但并不相等，因为时区不同
	fmt.Println(timeHangzhou.Equal(timeTokyo))
	timeTokyo = time.Date(2022, 4, 6, 22, 37, 48, 0, tokyo)
	// 看似差了一个小时，但因为时区不同
	fmt.Println(timeHangzhou.Equal(timeTokyo))
	// 本质上是看绝对时间，即时间戳
	fmt.Println(timeHangzhou.Unix(), timeTokyo.Unix())
}

// 时间操作
func timeTool3() {
	now := time.Now()
	// 当前时间加一天
	add := now.Add(24 * time.Hour)
	fmt.Println("+1 天 ", add)
	// 注：time.Hour 是 int64，那么编译器会自动推断 24 也是 int64，如果是下面的情况需要手动转
	dayUnit := 24
	add = now.Add(time.Duration(dayUnit) * time.Hour)
	fmt.Println("+1 天 ", add)
	// 求两个时间时间差
	sub := now.Sub(add)
	fmt.Println(sub.Seconds())
	// 时间前后判断
	fmt.Println(now.Before(add), now.After(add))
}

// 定时器
func timeTool4() {
	// 本质上是个通道
	for i := range time.Tick(time.Second) {
		fmt.Println(i)
	}
}

// DATE_FORMATTER 奇葩的时间格式化
var dataFormatter = map[string]string{
	"yyyy": "2006",
	"MM":   "01",
	"dd":   "02",
	"HH":   "15",
	"mm":   "04",
	"ss":   "05",
	"SSS":  "000",
}

func transform(format string) string {
	for k, v := range dataFormatter {
		format = strings.Replace(format, k, v, 1)
	}
	return format
}

func timeTool5() {
	now := time.Now()
	fmt.Println(now.Format(transform("yyyy-MM-dd HH:mm:ss.SSS")))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
}
