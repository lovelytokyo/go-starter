package main

import (
	"time"
	"fmt"
)

var timeNowFunc = time.Now

func main() {
	date := "2016-07-24"
	t, _ := time.Parse("2006-01-02", date)
	threeDaysAgo := timeNowFunc().Add(3 * -24 * time.Hour)

	fmt.Println(t.Unix()) // 指定日付
	fmt.Println(threeDaysAgo.Unix()) // ３日前日付
	fmt.Println(t.Unix() > threeDaysAgo.Unix())
}
