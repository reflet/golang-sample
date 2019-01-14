package main

import (
    "fmt"
    "strconv"
)

func main() {
    Variable()
    Variable1()
    Variable2()
    Variable3()
    Global()
    LocalVariable()
}

/**
 * 変数の宣言
 */
func Variable() {
    var (
        x int       // 型宣言（型のゼロ値が自動で挿入される）
        y int = 1   // 型宣言＋初期値
        z = 2       // 処置機＋初期値（型宣言を省略）
    )
    fmt.Println(x, y, z)
}

/**
 * 変数の宣言(複数)
 */
func Variable1() {
    var x1, x2 int
    var y1, y2 = 1, 2
    var z1, z2 = 1, "a"
    fmt.Println(x1, x2)
    fmt.Println(y1, y2)
    fmt.Println(z1, z2)
}

/**
 * 変数の宣言(複数 - 関数の戻り値あり)
 */
func Variable2() {
    var i1, err1 = strconv.Atoi("10")
    var _, err2 = strconv.Atoi("10") // i2は使わないので「 _ 」(ブランク識別子)で宣言する
    fmt.Println(i1, err1)
    fmt.Println(err2)
}

/**
 * 変数を宣言(:=)
 * ※ varを使って宣言した場合と内容は同じ
 * ※ 初期値があり、型宣言を省略できる場合に利用可能
 * ※ if文, for文などの初期化では、「:=」のみが有効（varキーワードは使えない）
 */
func Variable3() {
    x := 1
    y, z := 2, 3
    fmt.Println(x, y, z)
}

// グローバル変数の宣言
var n = 1

func Global() {
    fmt.Println(n)
    add()
    fmt.Println(n)
}

/**
 * グローバル変数の更新
 */
func add() {
    n += 1
}

/**
 * ローカル変数の宣言
 * ※ スコープ: {}(波括弧)内でのみ利用可能
 */
func LocalVariable() {
    // x はこの関数内であればどこでも利用可能（if文の中もOK）
    x := 1
    if x == 1 {
        // y はこのif文の中でだけ利用可能（if文の外ではスコープがない）
        y := 2
        fmt.Println(x, y)
    }
}
