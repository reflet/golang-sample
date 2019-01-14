package main

import "fmt"

type point struct {
    x, y int
}

/**
 * new組み込み関数
 * ※ 新規にメモリを割り当てる関数
 * ※ 複合型を指定したnew組み込み関数 = アドレス演算子「＆」を適用した複合リテラル
 *    （どちらも新規にメモリを割り当ててアドレスを取得する）
 */
func main() {
    // new組み込み関数の使用（数値型）
    n := new(int)       // nは、int型のポインタ型
    fmt.Println(n, *n)

    // new組み込み関数の使用（構造体型）
    p1 := new(point)    // p1は、point型のポインタ型
    fmt.Println(*p1)

    // 複合リテラルのアドレス取得（new組み込み関数と同じ意味）
    p2 := &point{}      // p2は、point型のポイント型
    fmt.Println(*p2)
}
