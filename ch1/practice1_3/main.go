package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "

	}
	fmt.Println(s)
	fmt.Println(time.Since(start).Milliseconds())
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start).Milliseconds())
}
