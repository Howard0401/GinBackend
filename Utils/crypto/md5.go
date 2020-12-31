package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

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
