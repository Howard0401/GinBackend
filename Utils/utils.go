package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"
)

//判斷query的參數是否會分頁，因此列表一定要分頁
func Page(Limit, Page int) (limit, offset int) {
	if limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

//給gorm用的排序
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

//時間轉換
const TimeLayout = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}

func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		log.Printf("transform time format err:%v", err)
	}
	return result.In(Local).Format(TimeLayout)
}

//加密
func Md5(str string) string {
	w := md5.New()
	_, err := io.WriteString(w, str)
	if err != nil {
		log.Printf("md5 err :%v", err)
	}
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
