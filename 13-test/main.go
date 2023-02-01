package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
)

func main() {
	appid := "20230201001545974"
	secretKey := "M0MUrnjoQtmWwnB4a877"
	myurl := "/api/trans/vip/translate"
	fromLang := "auto"
	toLang := "zh"
	salt := rand.Intn(32768) + 32768
	q := "apple"
	sign := appid + q + string(salt) + secretKey
	fmt.Println(sign)

	h := md5.New()
	io.WriteString(h, sign)
	sign = fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(sign)
	myurl = myurl + "?appid=" + appid + "&q=" + q + "&from=" + fromLang + "&to=" + toLang + "&salt=" + strconv.Itoa(salt) + "&sign=" + sign
	fmt.Println(myurl)
}
