package main

import "fmt"

func init() {
	fmt.Printf("最初に呼ばれる\n")
}

func hoge() {
	fmt.Printf("hoge!\n")
}

func main() {
	hoge()
	fmt.Printf("Hello world\n")
}
