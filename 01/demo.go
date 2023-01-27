package main

import (
	"fmt"
	"time"
)

func helloworld() {
	fmt.Println("hello,wordld!")
	s := fmt.Sprintln("hello,world")
	fmt.Println(s)
}
func vartest() {
	var a = "initial"
	var b, c int = 1, 2
	var d = true
	var e float64
	f := float32(3.0)
	g := a + "start"
	fmt.Println(a, b, c, d, e, f, g)
}
func constTest() {
	const s = "const var"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i)
	const (
		a = iota
		b
		c
	)
	println(a, b, c)
}
func ifTest() {
	i := 0
	for {
		fmt.Println("loop", i)
		i += 1
		if i == 100 {
			break
		}
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
	for i > 1 {
		println(i)
		i -= 1
	}
}
func switchTest() {
	t := time.Now()
	switch {
	case t.Hour() > 12:
		println("下午")
	case t.Hour() < 12:
		println("上午")
	}
}

// var a string
// a="hello,world"
// a:="hello,world" equals
func arrayTest() {
	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}
	var c [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c[i][j] = i - j
		}
	}
	var d [6][6]bool
	fmt.Println(a, b, c, d, len(d), a[0])
}
func sliceTest() {
	var c = make([]string, 3)
	c[0] = "h"
	c[2] = "l"
	//print(c)
	fmt.Println(c)
	c = append(c, "l")
	fmt.Println(len(c))
	c = append(c, "o")
	fmt.Println(len(c))
	d := make([]string, 6)
	copy(d, c)
	fmt.Println(c, d)
	fmt.Println(d[0:5])
	fmt.Println(len(d))
}
func mapTest() {
	c := make(map[string]string)
	c["a"] = "i"
	c["b"] = "am"
	fmt.Println(c)
	d, e := c["undefined"]
	fmt.Println(d, e)
	delete(c, "a")
	m2 := map[string]int{"one": 1}
	var m3 = map[string]int{"two": 2, "three": 3}
	fmt.Println(m2, m3)
	fmt.Println(d)
}
func main() {
	//helloworld()
	//vartest()
	//constTest()
	//ifTest()
	//switchTest()
	//arrayTest()
	//sliceTest()
	mapTest()
}
