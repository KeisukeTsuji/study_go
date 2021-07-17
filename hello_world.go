package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("最初に呼ばれる")
}

func hoge() {
	fmt.Println("hoge!")
}

func printlnTimeNow() {
	// https://pkg.go.dev/time@go1.16.6#Now
	fmt.Println(time.Now())
}

func printlnUserCurrent() {
	// https://pkg.go.dev/os/user@go1.16.6#Current
	fmt.Println(user.Current())
}

func main() {
	hoge()
	printlnTimeNow()
	printlnUserCurrent()
	fmt.Println("Hello world")
}
