package main

import (
    "fmt"
    "unsafe"
)

func main() {
    Constant()
    Iota()
    Expression()
}

/**
 * 定数の宣言
 */
func Constant() {
    const (
        c1 uint = 1  // uint型を持つ（型の記述があるため）
        c2 = c1      // uint型を持つ（c1がuint型を持つため）
        c3 = 1       // 型を持たない
    )

    /**
     * / 定数名のみの記述の場合、１つ上の行と同じリテラル（または定数式）が設定される
     * cont (
     *     u1 uint = 1 // 1
     *     u2          // u1と同じリテラルを設定
     *     u3          // u1と同じリテラルを設定
     * )
     */

    var (
        // 定数c3は型を持たないため、どの数値型にも使用可能
        v1 int = c3
        v2 uint = c3

        // 定数c1, c2はuint型の定数のため、int型の変数には使えない
        // v3 int = c1
        // v3 int = c2
    )

    fmt.Println(c1, c2, c3)
    fmt.Println(v1, v2)
}

/**
 * iota(イオタ)
 * ※ enumのように連続した値の定数を作成したい場合に利用する
 * ※ iotaで作成された定数は、整数リテラルや文字列リテラルのように型を持たない
 */
func Iota() {
    const (
        a0 = iota           // 0
        a1 = iota           // 1
        a2 = 9              // 9
        a3 = iota           // 3 (iotaの値は丸カッコ内での行数で決まる)
        a4, a5 = iota, iota // a4, a5ともに値4
        a6 = iota * iota    // 25 (5 * 5の結果)
    )
    fmt.Println(a0, a1, a2, a3, a4, a5, a6)

    const (
        b0 = iota           // 0
        b1                  // 1 (値を省略した場合は、1行上の式「iota」を使用する
        b2                  // 2
        _                   // ブランク識別子によるスキップ (式自体は、iota)
        b3                  // 4
    )
    fmt.Println(b0, b1, b2, b3)

    const (
        c0 = 1 << iota      // 1
        c1                  // 2 (値を省略した場合は、1行上の式「1 << iota」を使用する
        c2                  // 4
        c3                  // 8
        c4                  // 16
    )
    fmt.Println(c0, c1, c2, c3, c4)

    const d0 = iota // 0 (丸カッコを使わない限り、iotaの値は0)
    const d1 = iota // 0 (丸カッコを使わない限り、iotaの値は0)
    fmt.Println(d0, d1)
}

/**
 * 定数式
 */
func Expression() {
    const (
        intLit       = 123                     // 数値リテラル
        strLit       = "ABC"                   // 文字列リテラル
        intConst     = -intLit                 // 定数同士の演算(整数)
        strConst     = strLit + "DEF"          // 定数同士の演算(文字列)
        floatConst   = float64(intLit + 456)   // 型変換
        complexConst = complex(1.1, 2.2) // complex組込み関数
        realConst    = real(complexConst)      // real組込み関数
        imagConst    = imag(complexConst)      // imag組込み関数
    )

    // len組込み関数、cap組込み関数で使用する配列
    var a = [3]int{1, 2, 3}

    // unsafeの組込み関数で使用する構造体変数
    var s = struct{ x int64; y float64 }{}

    const (
        lenConst = len(a)                    // len組込み関数
        capConst = cap(a)                    // cap組込み関数
        alignofConst = unsafe.Alignof(s)     // unsafe.Alignof組込み関数
        offsetofConst = unsafe.Offsetof(s.y) // unsafe.Offsetof組込み関数
        sizeofConst = unsafe.Sizeof(s)       // unsetSizeof組込み関数
    )

    fmt.Println(intConst, strConst, floatConst)
    fmt.Println(complexConst, realConst, imagConst)
    fmt.Println(lenConst, capConst)
    fmt.Println(alignofConst, offsetofConst, sizeofConst)
}
