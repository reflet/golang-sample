package workopt

import "fmt"

/**
 * 演算子
 */
func Operator() {
    var (
        // 四則演算
        v1 = 2 + 3        // 5 (通常の加算)

        // 四則演算によるオーバーフロー
        u1 uint8 = 255    // 255は、uint8型の最大値
        v2 = u1 + 1       // オーバーフローのため、値は「0」になる

        u2 int8 = 127     // 127は、int8型の最大値
        v3 = u2 + 1       // オーバーフローのため、値は「-128」になる

        // ビット演算
        u3 uint8 = 0xA3   //「0xA3」を2進数で表記すると「10100011」、10進数では「163」
        u4 uint8 = 0x53   //「0x53」を2進数で表記すると「01010011」、10進数では「83」
        v4 = u3 & u4      // 変数v4を2進数で表記すると、「00000011」、10進数では「3」
        v5 = u3 | u4      // 変数v4を2進数で表記すると、「11110011」、10進数では「243」
        v6 = u3 ^ u4      // 変数v4を2進数で表記すると、「11110000」、10進数では「240」
        v7 = u3 &^ u4     // 変数v4を2進数で表記すると、「10100000」、10進数では「160」
        v8 = ^u3          // 変数v4を2進数で表記すると、「01011100」、10進数では「92」

        // シフト演算子（符号なし整数）
        v9 = u3 << 4      // 変数v9を2進数で表記すると、「00110000」、10進数では「48」
        v10 = u3 >> 4     // 変数v10を2進数で表記すると、「00001010」、10進数では「10」

        // シフト演算子（符号付き整数）
        i1 int8 = -0x5D   // int8型の「-0x50」を２進数で表記すると「10100011」
        v11 = i1 << 4     // 変数v11を2進数で表記すると、「00110000」、10進数では「48」
        v12 = i1 >> 4     // 変数v12を2進数で表記すると、「11111010」、10進数では「-6」

        // 型エラーの例（int型とuint型の演算）
        // i int = 1
        // u uint = 1
        // v = i + u
    )

    fmt.Println(v1, v2, v3)
    fmt.Println(v4, v5, v6, v7, v8)
    fmt.Println(v9, v10)
    fmt.Println(v11, v12)
}

/**
 * 比較演算子
 */
func ComparisonOperator() {
    // 整数型の比較演算子
    fmt.Println(1 == 1, 1 != 2)
    fmt.Println(1 <= 2, 1 < 2)
    fmt.Println(2 >= 1, 2 > 1)

    // 文字列型の比較演算子
    fmt.Println("ABC" == "ABC", "A" < "B", "ABC" < "ACB")
}

/**
 * 論理演算子
 */
func LogicalOperator() {
    // 左側(trueFunc)のみ評価される。右側は評価されない
    fmt.Println(trueFunc() || falseFunc())

    // 左側(falseFunc)のみ評価される。右型は評価されない
    fmt.Println(falseFunc() && trueFunc())

    // 両方とも評価される
    fmt.Println(falseFunc() || trueFunc())
    fmt.Println(trueFunc() && trueFunc())
}
func trueFunc() bool {
    fmt.Print("[trueFunc]")
    return true
}
func falseFunc() bool {
    fmt.Print("[falseFunc]")
    return false
}

/**
 * インクリメント/デクリメント
 */
func IncrementDecrement() {
    i := 1

    // インクリメント（変数iに1を加算）
    // ※ ++iは存在しない
    i++
    fmt.Println(i)

    // デクリメント（変数iから1を減算）
    // ※ --iは存在しない
    i--
    fmt.Println(i)

    // 下記のような書き方はできない
    // ※ 式ではないので代入文には使用できない
    // a[i++]    // 配列のアクセス直後のインクリメント
    // b = i++   // 代入直後のインクリメント
}

/**
 * 複素数
 */
func ComplexCalc() {
    var (
        f32r float32 = 1.1
        f32i float32 = 2.2
        f64r float64 = 1.1
        f64i float64 = 2.2
    )

    // float32型の引数から、complex64型；の複素数を取得
    var c64 complex64 = complex(f32r, f32i)
    fmt.Println("c64: ", c64)

    // float64型の引数から、complex128型；の複素数を取得
    var c128 complex128 = complex(f64r, f64i)
    fmt.Println("c128: ", c128)

    // real関数、imag関数
    // ※ complex64型の引数から、float32型の実数と虚数を取得
    f32r = real(c64)
    f32i = imag(c64)
    fmt.Println("f32r: ", f32r)
    fmt.Println("f32i: ", f32i)

    // real関数、imag関数
    // ※ complex128型の引数から、float64型の実数と虚数を取得
    f64r = real(c128)
    f64i = imag(c128)
    fmt.Println("f64r: ", f64r)
    fmt.Println("f64i: ", f64i)
}

/**
 * 文字列の操作
 */
func StrOperation() {
    s1 := "abc"
    s2 := "あいうえお"

    // 文字列「abc」と「あいうえお」の結合
    s3 := s1 + s2
    fmt.Println(s3)       // abcあいうえお
    fmt.Println(len(s3))  // 18 (文字列「abcあいうえお」のバイト長)
}
