package main

import (
	"fmt"
	"sort"
)

func minInt(a []int) int {
	sort.Ints(sort.IntSlice(a))
	return a[0]
}

func sum(s map[string]int) (n int) {
	for _, t := range s {
		n += t
	}
	return
}

func main() {
	f1 := 1.11
	f2 := int(f1)
	fmt.Println(f2)

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	fmt.Println(minInt(l))

	s := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	fmt.Println(sum(s))
}
