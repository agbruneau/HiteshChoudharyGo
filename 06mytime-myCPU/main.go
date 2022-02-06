package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("Monday 01-02-2006 15:04:05"))
	fmt.Println("Number of CPU on host: ", runtime.NumCPU())
}
