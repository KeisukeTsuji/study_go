/*
Goで学ぶポインタとアドレス
https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
*/

package main

import "fmt"

func main() {
	var n int = 100

	fmt.Println(n)

	// アドレス
	fmt.Println(&n)

	// intのポイント型
	var p *int = &n

	fmt.Println(p)

	// アドレスの指しているメモリの中身
	fmt.Println(*p)

}
