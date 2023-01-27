package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	serectNum := rand.Intn(maxNum)
	fmt.Println(serectNum)
	var s int
	var count int
	fmt.Println("请输入0-100的正整数")
	for {
		_, err := fmt.Scanf("%d", &s)
		if err != nil {
			fmt.Println("输入不合法，请重试", err)
			continue
		}
		count++
		fmt.Println("您猜的数为；", s)
		if s < serectNum {
			fmt.Println("您猜小了，请再次尝试吧")
		} else if s > serectNum {
			fmt.Println("您猜大了，请再次尝试吧")
		} else {
			fmt.Printf("您猜对了，总共用了%d次，恭喜您！\n", count)
			break
		}
	}

}
