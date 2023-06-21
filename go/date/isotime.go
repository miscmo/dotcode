package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 获取自然周的起始时间和结束时间
	year, week := now.ISOWeek()
	weekStart := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, (week-1)*7-int(now.Weekday()))
	weekEnd := weekStart.AddDate(0, 0, 7)

	// 获取自然月的起始时间和结束时间
	year, month, _ := now.Date()
	monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// 将时间格式化为 "2022-01-23" 的形式
	weekStartStr := weekStart.Format("2006-01-02")
	weekEndStr := weekEnd.Format("2006-01-02")
	monthStartStr := monthStart.Format("2006-01-02")
	monthEndStr := monthEnd.Format("2006-01-02")

	fmt.Println("Week start:", weekStartStr)
	fmt.Println("Week end:", weekEndStr)
	fmt.Println("Month start:", monthStartStr)
	fmt.Println("Month end:", monthEndStr)

	weekDS, weekDE := GetWeekDay()
	fmt.Println("week start: %s, end: %s", weekDS, weekDE)

	monthDS, monthDE := GetMonthDay()
	fmt.Println("month start: %s, end: %s", monthDS, monthDE)
}


func GetMonthDay() (string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02"), time.Unix(l, 0).Format("2006-01-02")
}

func GetWeekDay() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()

	return time.Unix(f, 0).Format("2006-01-02"), time.Unix(l, 0).Format("2006-01-02")
}