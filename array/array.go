package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b = [2]int{100, 200}
	fmt.Println(b)

	// 型で配列のlengthが決まっているので追加するとエラーになる。sliceなら可能。
	// b.append(300)

	n := []int{1, 2, 3, 4, 5, 6}
	n = append(n, 100)
	fmt.Println(n)
}
