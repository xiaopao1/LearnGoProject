package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	a := userInfo{"wang", 18, []string{"swim", "dance"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	println(buf)
	println(string(buf))
	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	println(string(buf))
	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", b)

}
