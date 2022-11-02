package main

import (
	"fmt"
	"os"
)

func main() {
	var s, step string
	// 1是编译后的缓存文件
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
		fmt.Printf("index: %v, value: %v\n", i, os.Args[i])
	}
	fmt.Println(s)
}
