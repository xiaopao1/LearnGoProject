package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	pi := 3.14159261542912309
	fmt.Printf("%T : %v\n", pi, pi)
	fmt.Println(math.MaxFloat64)
	var f1 float32 = 10000018
	var f2 float32 = 100000018 //float32的精度是小数点后7未，float64是15位好像
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f2 + 1)
	i, err := strconv.ParseInt("0xffff", 0, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %v \n", i, i)
	i1, err := strconv.ParseInt("012", 0, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %v \n", i1, i1)
	i2, err := strconv.ParseInt("0b101111", 0, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %v \n", i2, i2)
	fmt.Printf("%z \n", i2)
	v1, err := strconv.Atoi("1234")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %v\n", v1, v1)
	fv1, err := strconv.ParseFloat("3.1234567890123456789", 64)
	fmt.Printf("%z \n", fv1)
	fv2, err := strconv.ParseFloat("3.1234567890123456789", 32)
	fmt.Printf("%z \n", fv2)

}
