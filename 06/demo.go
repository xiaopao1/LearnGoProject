package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	t1 := time.Date(2002, 3, 12, 12, 12, 12, 0, time.UTC)
	t2 := time.Date(2003, 4, 12, 12, 12, 12, 0, time.UTC)
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())
	fmt.Println(t1.Format("2006:01:02 15-04-05"))
	fmt.Println(t2)
	diff := t1.Sub(t2)
	fmt.Println(diff)

	println(t.Unix())
	t3, err := time.Parse("2006:01:02 15-04-05", "2002:03:12 12-12-12")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3)
	fmt.Println(t3 == t1)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now())
		fmt.Println(time.Now().Unix())
	}
}
