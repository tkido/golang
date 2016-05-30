package main

import (
	"fmt"
	"math"
)

/*Sqrt is exproted function */
func Sqrt(x float64) float64 {
	n := x
	for i := 0; i < 5; i++ {
		n = n - ((n*n - x) / (2 * n))
	}
	return n
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))

	fmt.Println("ということで今日はAtomテキストエディタに表示している日本語フォントの設定を簡単に変更する方法をお伝えしますね。")
}
