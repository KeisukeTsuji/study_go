package main

import "fmt"

func main() {
	f1 := 1.11
	f2 := int(f1)
	fmt.Println(f2)

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
