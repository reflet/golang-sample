package main

import "fmt"

func main() {
    // 11.1.1 配列
    example11()

    // 11.1.6 配列リテラルとスライスリテラルの省略記法
    example16()

    // 11.1.7 配列とスライスのfor文
    example17()
}

/**
 * -------------------------------------------------------
 * 11.1.1 配列
 * -------------------------------------------------------
 * ※ 固定長 (後から追加NG)
 * ※ 値型: 変数ごとに必要な長さを格納できるメモリを割り当てる
 * ※「==」演算子: 使用可能 (型は各変数ごとに異なる)
 */
func printArrayInt(a [3]int) {
    fmt.Println("array:", a[0], a[1], a[2])
}

func example11() {
    var array1 [3]int              // すべてゼロ値がセットされる

    // 初期値に配列リテラルを指定した変数宣言
    array2 := [3]int{20, 21}       // 長さを指定する配列リテラル（３番めの要素はゼロ値「0」がセットされる）
    array3 := [...]int{30, 31, 32} // 長さを指定しない配列リテラル
    array4 := [3]int{1: 41}        // インデックスを指定した配列リテラル

    printArrayInt(array1)
    printArrayInt(array2)
    printArrayInt(array3)
    printArrayInt(array4)
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
