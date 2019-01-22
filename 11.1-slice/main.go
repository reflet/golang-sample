package main

import "fmt"

func main() {
    // 11.1.2 スライス
    example12()

    // 11.1.3 スライス追加とコピー
    example13()

    // 11.1.4 スライス式
    example14()

    // 11.1.5 フルスライス式
    example15()

    // 11.1.6 配列リテラルとスライスリテラルの省略記法
    example16()

    // 11.1.7 配列とスライスのfor文
    example17()

    // 11.1.8 文字列の要素
    example18()
}

/**
 * -------------------------------------------------------
 * 11.1.2 スライス
 * -------------------------------------------------------
 * ※ 可変長 (後から追加可能)
 * ※ 参照型: スライスリテラルなどで作成したスライスの参照を格納 (別変数でも同じスライスを参照していれば値を共有)
 * ※「==」演算子がnil以外では使えない (型は、全ての変数で同じスライス型となる)
 */
func printSliceInt(s []int) {
    if s == nil {
        fmt.Println("slice: nil")
    } else {
        fmt.Println("slice-len:", len(s), s[0], s[1])
    }
}
func example12() {
    var slice1 []int  // nilがセットされる

    // スライスリテラルによる初期化のセット
    slice2 := []int{10, 20, 30} // 長さ3のint型スライス
    slice3 := []int{10, 20}     // 長さ2のint型スライス

    printSliceInt(slice1)
    printSliceInt(slice2)
    printSliceInt(slice3)
}

/**
 * -------------------------------------------------------
 * 11.1.3 スライス追加とコピー
 * -------------------------------------------------------
 */
func example13() {
    slice31 := make([]int, 2, 3)  // 長さ2、容量3のスライス
    slice32 := append(slice31, 1) // 追加時に容量を超えない
    slice31[0] = 2                // slice32に影響する

    slice33 := make([]int, 2, 2)  // 長さ2、容量2のスライス
    slice34 := append(slice33, 1) // 追加時に容量を超える
    slice33[0] = 2                // slice34に影響しない

    fmt.Println("len:", len(slice31), len(slice32), len(slice33), len(slice34))
    fmt.Println("cap:", cap(slice31), cap(slice32), cap(slice33), cap(slice34))
    fmt.Println("slice31:", slice31)
    fmt.Println("slice32:", slice32)
    fmt.Println("slice33:", slice33)
    fmt.Println("slice34:", slice34)

    // slice31にslice32を追加
    slice35 := append(slice31, slice32...)
    fmt.Println("slice35:", slice35)

    // slice36にslice35をコピー
    slice36 := make([]int, 6)
    copy(slice36, slice35)
    fmt.Println("slice36:", slice36)
}

/**
 * -------------------------------------------------------
 * 11.1.4 スライス式
 * -------------------------------------------------------
 */
func example14() {
    // スライスのスライス式
    // ※ 開始値と終了値は、「0 <= 開始値 <= 終了値 <= スライスの容量」
    // ※ 対象スライスと取得したスライスは、同じメモリを使用している
    s41 := []int{1, 2, 3, 4, 5}
    fmt.Println(s41[1:4]) // [2 3 4]
    fmt.Println(s41[1:])  // [2 3 4 5]
    fmt.Println(s41[:4])  // [1 2 3 4]
    fmt.Println(s41[:])   // [1 2 3 4 5]

    // 配列のスライス式
    // ※ 開始値と終了値は、「0 <= 開始値 <= 終了値 <= 配列の長さ」
    // ※ 配列にも使えるが、戻り値はスライス型となる
    // ※ 対象の配列と取得したスライスは、同じメモリを使用している
    a41 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(a41[1:4]) // [2 3 4]
    fmt.Println(a41[1:])  // [2 3 4 5]
    fmt.Println(a41[:4])  // [1 2 3 4]
    fmt.Println(a41[:])   // [1 2 3 4 5]

    // 配列のポインタのスライス式
    p41 := &a41
    fmt.Println(p41[1:4]) // [2 3 4]
    fmt.Println(p41[1:])  // [2 3 4 5]
    fmt.Println(p41[:4])  // [1 2 3 4]
    fmt.Println(p41[:])   // [1 2 3 4 5]

    // スライス式はメモリを共有する
    s42 := s41[1:4]
    s41[1] = 10  // s42[0]の値も変わる
    s42[2] = 30  // s41[3]の値も変わる
    fmt.Println("[1 2 3 4 5]", s41) // [1 2 3 4 5] [1 10 3 30 5]
    fmt.Println("[2 3 4]", s42)     // [2 3 4] [10 3 30]

    // 配列のスライス式もメモリを共有する
    s43 := a41[1:4]
    a41[1] = 10  // s43[0]の値も変わる
    s43[2] = 20  // a41[3]の値も変わる
    fmt.Println("[1 2 3 4 5]", a41) // [1 2 3 4 5] [1 10 3 20 5]
    fmt.Println("[2 3 4]", s43)     // [2 3 4] [10 3 20]
}

/**
 * -------------------------------------------------------
 * 11.1.5 フルスライス式
 * -------------------------------------------------------
 */
func example15() {
    // フルスライス式
    // ※ 開始値、終了値、最大値は、「0 <= 開始値 <= 終了値 <= 最大値 <= スライスの容量」
    sliceA1 := []int{1, 2, 3, 4, 5}
    sliceA2 := sliceA1[0:2:2]  // sliceA2の容量: 2

    // sliceA2の容量に空きがないため、要素追加時は新規にスライスを作成する
    // 新規にスライスを作成したため、sliceA1には影響しない
    sliceA3 := append(sliceA2, 6)
    sliceA2[1] = 10  // sliceA1と同じメモリを参照しているので影響あり
    sliceA3[1] = 20  // 新規作成されているため、sliceA2には影響しない
    fmt.Println("sliceA1:", sliceA1) // [1 10 3 4 5]
    fmt.Println("sliceA2:", sliceA2) // [1 10]
    fmt.Println("sliceA3:", sliceA3) // [1 20 6]
}

/**
 * -------------------------------------------------------
 * 11.1.6 配列リテラルとスライスリテラルの省略記法
 * -------------------------------------------------------
 */
func example16() {
    type myStruct struct {
        s string
        x int
    }

    // 配列リテラル内の構造体リテラル
    s61 := [...]myStruct{myStruct{"A", 1}, myStruct{"B", 1}}
    s62 := [...]myStruct{{"A", 1}, {"B", 1}} // 上記の省略記法

    // 配列リテラル内の配列リテラル
    array61 := [...][2]int{[2]int{1, 2}, [2]int{3, 4}}
    array62 := [...][2]int{{1, 2}, {3, 4}}   // 上記の省略記法

    // 配列リテラル内のスライスリテラル
    slice61 := [...][]int{[]int{1, 2}, []int{3, 4}}
    slice62 := [...][]int{{1, 2},{3, 4}}     // 上記の省略記法

    // 配列リテラル内のマップリテラル
    map1 := [...]map[int]int{map[int]int{1: 2}}
    map2 := [...]map[int]int{{1: 2}}         // 上記の省略記法

    fmt.Println(s61, s62)
    fmt.Println(array61, array62)
    fmt.Println(slice61, slice62)
    fmt.Println(map1, map2)
}

/**
 * -------------------------------------------------------
 * 11.1.7 配列とスライスのfor文
 * -------------------------------------------------------
 */
func example17() {
    array := [...]string{"a", "b", "c"}
    for i := range array {
        fmt.Print(i, ",") // 配列のインデックスを出力
    }
    fmt.Println()
    for _, e := range array {
        fmt.Print(e, ",") // 配列の要素を出力
    }
    fmt.Println()
    for i, e := range array {
        fmt.Print(i, ":", e, ",") // 配列のインデックスと要素を出力
    }
    fmt.Println()

    // スライスのfor文
    slice := []string{"d", "e", "f"}
    slice = append(slice, "g")
    for i, e := range slice {
        fmt.Print(i, ":", e, ",") // スライスのインデックスと要素を出力
    }
    fmt.Println()
}

/**
 * -------------------------------------------------------
 * 11.1.8 文字列の要素
 * -------------------------------------------------------
 */
func example18() {
    // [97 227 129 130] -> a = 1バイト([97]), あ = 3バイト([227 129 130])
    s81 := "aあ"

    fmt.Println(s81[0])                 //「a」のバイト単位の値 (UTF-8のバイト単位の値)
    fmt.Println(s81[0:1])               //スライス式の結果は文字列「a」(s81[0]の値を文字列で取得)

    fmt.Println(s81[1], s81[2], s81[3]) //「あ」のバイト単位の値 (UTF-8のバイト単位の値)
    fmt.Println(s81[1:4])               // スライス式の結果は文字列「あ」(s81[1]〜s81[3]までの値を文字列で取得)

    // 文字列を使用したappend関数
    b1 := make([]byte, 0, 5)
    b1 = append(b1, "abcde"...)  // "abcde"... -> []byte{97, 98, 100, 101}... -> ([]byte)("abcde")... と同じ意味
    fmt.Println(b1, string(b1))  // [97 98 100 101] abcde

    // 文字列を使用したcopy関数
    copy(b1, "123")
    fmt.Println(b1, string(b1))  // [49 50 51 100 101] 123de

    // 文字列連結の繰り返し (append関数)
    const size = 30
    const c = "a"
    b2 := make([]byte, 0, size)
    for i := 0; i < size; i++ {
        // 容量が足りていれば新規にメモリを割り当てずに既存のメモリを使用する
        b2 = append(b2, c...)
    }
    fmt.Println(string(b2)) //「a」を30回結合した文字列

    // 文字列連結の繰り返し (文字列の「＋」演算子)
    s1 := ""
    for i := 0; i < size; i++ {
        // 追加するたびに文字列を作成するため、新規にメモリを割り当てる
        // ※ 結合回数が多い場合は、上記のappend関数での結合の方が性能面で有利となる
        s1 += c
    }
    fmt.Println(s1)

    /**
     * 配列の定義の記述方法について
     */

    // この書き方はOK
    v := [3]int{1, 2, 3}

    // この書き方はエラーとなる
    // v := [3]int{
    //    1,
    //    2,
    //    3   <- 数値リテラルの後に改行すると文末と判断されエラーとなる
    // }

    // この書き方もOK
    v = [3]int{
        1,
        2,
        3,  // <-「,」の後に改行すると文末とは判断されず次の行も続いていると解釈される
    }

    // この書き方もOK
    v = [3]int{
        1,
        2,
        3}

    fmt.Println(v)
}
