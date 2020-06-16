package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	fmt.Println("请输入你的名字")
	for idx,arg := range os.Args[1:] {
		s += sep + os.Args[1]
		sep = " "
		fmt.Println("参数"+strconv.Itoa(idx)+":",arg)
	}
	fmt.Println(s)
}
