package main

import "fmt"

func rangeTest() {
	for i, c := range "helloworld" {
		fmt.Println(i, c)
	}
	m1 := make(map[int]float32)
	m1[1] = 1.0
	m1[2] = 2.0
	m1[3] = 3.0
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println("map m1 key is", k)
	}
	for _, v := range m1 {
		fmt.Println("map m1 value is", v)
	}
}
func add(a int, b int) int {
	return a + b
}
func exist(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}
func funcTest() {
	//s := make([]int, 9)
	k, v := exist(map[string]string{"a": "A", "b": "B"}, "c")
	fmt.Println(k, v)
}
func addn(n int) {
	n += 2
}
func addnPtr(n *int) {
	*n += 2
}
func pointTest() {
	a := 2
	addn(a)
	fmt.Println(a)
	addnPtr(&a)
	fmt.Println(a)
}

type user struct {
	name     string
	password string
}

func checkPassword(u user, password string) bool {
	return u.password == password
}
func checkPassword1(u *user, password string) bool {
	return u.password == password
}
func (u *user) resetPassword(password string) {
	u.password = password
}
func (u user) resetPassword1(password string) {
	u.password = password
}
func structTest() {
	z := user{name: "ma", password: "jack"}
	a := user{"li", "123."}
	b := user{name: "wang"}
	var d user
	d.name = "zhang"
	d.password = "zadas"
	fmt.Println(z, a, b, d)
	fmt.Println(checkPassword(a, "123"))
	fmt.Println(checkPassword1(&a, "123"))
	a.resetPassword1("123")
	fmt.Println(checkPassword1(&a, "123"))
	a.resetPassword("123")
	fmt.Println(checkPassword1(&a, "123"))

}
func main() {
	//rangeTest()
	//funcTest()
	//pointTest()
	structTest()
}
