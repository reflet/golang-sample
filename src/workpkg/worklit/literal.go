package worklit

import "fmt"

/**
 * リテラル (整数)
 */
func Literal() {
    var (
        // 整数リテラル
        n1 = 6789   // 10進数
        n2 = 04567  // 8進数
        n3 = 0xCDEF // 16進数

        // 浮動小数点リテラル
        f1 = 1.2
        f2 = 1.2e+3 // 指数表記

        // 虚数リテラル
        i1 = 1.2i     // 虚数のみの複素数
        i2 = 3 + 1.2i // 実数と虚数を持つ複素数

        // 文字（ルーン）リテラル
        r1 = 'a'
        r2 = 'あ'
        r3 = '\n'         // エスケープシーケンス
        r4 = '\141'       // 8進数Unicodeコードポイント
        r5 = '\u0061'     // 16進数Unicodeコードポイント (4桁)
        r6 = '\U00000061' // 16進数Unicodeコードポイント (8桁)
        r7 = '\x61'       // 16進数UTF-8 (2桁)

        // 下記のリテラルは、整数型でも受け取れる
        b1 byte = 1.0     // 小数点以下に数値がない
        b2 byte = 0i      // 虚数が0
    )

    fmt.Println(n1, n2, n3)
    fmt.Println(f1, f2)
    fmt.Println(i1, i2)
    fmt.Println(r1, r2, r3, r4, r5, r6, r7)
    fmt.Println(b1, b2)
}

/**
 * リテラル (文字列)
 */
func LiteralStr() {
    var (
        // interpreted (解釈有り) 文字列リテラル
        // ※ ダブルクォートで囲む
        s1 = "abc\nあいうえお\nかきくけこ"

        // raw (生) 文字列リテラル
        // ※ バッククォートで囲む
        s2 = `abc
あいうえお
かきくけこ`

        // エスケープシーケンスによる「あ」
        s3 = "\u3042"        // Unicodeコードポイント指定
        s4 = "\xE3\x81\x82"  // UTF-8指定 (3バイト)
    )

    fmt.Println(s1, s2, s3, s4)
}
