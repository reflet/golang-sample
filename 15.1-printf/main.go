package main

import "fmt"

/**
 * -------------------------------------------------------
 * 15.1 ツール・ライブラリ - 文字の出力
 * -------------------------------------------------------
 */
func main() {
    // 15.1.1 Printf関数
    example11()

    // 15.1.2 デフォルトの出力形式の変更 (Stringメソッド)
    example12()
}

/**
 * -------------------------------------------------------
 * 15.1.1 Printf関数
 * -------------------------------------------------------
 * func Printf(format string, a ...interface{}) (n int, errno os.Error)
 * formatについては、fmtパッケージの公式ドキュメントを参照する
 * https://golang.org/pkg/fmt
 */
func example11() {
    fmt.Printf("[%v]\n", "ABC")       // デフォルト形式
    fmt.Printf("[%#v]\n", "ABC")      // ダブルクォート付き
    fmt.Printf("[%+d]\n", 10)         // 符号付き
    fmt.Printf("[%4d]\n", 10)         // 左パディング(空白埋め)
    fmt.Printf("[%04d]\n", 10)        // 左パディング(0埋め)
    fmt.Printf("[%6.2f]\n", 1234.567) // 小数点以下の精度指定
}

/**
 * -------------------------------------------------------
 * 15.1.2 デフォルトの出力形式の変更 (Stringメソッド)
 * -------------------------------------------------------
 * ※ Print関数やPrintf関数の引数の文字列で使用される「%v」などは、それぞれの型に用意されたデフォルトの形式で出力される
 * ※ この出力形式を変更したい場合は、fmtパッケージのStringerインターフェースを実装する
 *
 * [Stringerインターフェース]
 *
 *   type Stringer interface {
 *       String() string
 *   }
 */
type MyString string

// MyString型のメソッドの宣言(String)
func (s MyString) String() string {
    return "***" + string(s) + "****"
}

func example12() {
    var s MyString = "ABC"

    // Print関数は、MyString型のStringメソッドの戻り値を出力する
    fmt.Println(s)
}
