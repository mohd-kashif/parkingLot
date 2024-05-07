package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	time.Sleep(5 * time.Second)
	c := time.Since(a).Seconds()
	fmt.Println(c)
}
