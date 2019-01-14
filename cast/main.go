package main

import "fmt"

/**
 * 型変換（キャスト）
 */
func main() {
    var (
        u8 uint8
        u16 uint16
        f64 float64
        r = 'あ'
        str string = "abcあいうえお"
        byteSlice []byte
        runeSlice []rune
    )

    // サイズの異なる型への変換
    u8 = uint8(u16)
    u16 = uint16(u8)
    fmt.Println(u8, u16)

    // 整数型と浮動小数点型の変換
    f64 = float64(u16)  // 整数型から浮動小数点型への変換
    u16 = uint16(f64)   // 浮動小数点型から整数型への変換
    fmt.Println(f64, u16)

    // 整数型（rune型）から文字列型への変換
    var s1 string = string(r)
    fmt.Println(s1)     //「あ」を出力

    // 文字列型からbyteのスライス型に変換
    byteSlice = []byte(str)
    fmt.Println(byteSlice)   // バイト文字列単位の出力（18個）

    // 文字列型からruneのスライス型に変換
    runeSlice = []rune(str)
    fmt.Println(runeSlice)   // マルチバイト文字列単位の出力（８個）

    var s2 string = string(byteSlice)
    fmt.Println(s2)          //「abcあいうえお」を出力

    var s3 string = string(runeSlice)
    fmt.Println(s3)          //「abcあいうえお」を出力
}
