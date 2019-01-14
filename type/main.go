package main

import "fmt"

type myInt32 int32

type (
    myInt int
    myString string
)

func main() {
    TypeDeclaration()
    NamedType()
}

/**
 * 独自の型の宣言
 */
func TypeDeclaration() {
    // 型の宣言で作成したmyInt32型を使用
    var n1 myInt32 = 1
    var n2 myInt32 = 2
    fmt.Println(n1 + n2)

    // 異なる型の間の演算はできない
    var i1 int32 = 1
    // fmt.Println(n1 + i1)
    fmt.Println(n1 + myInt32(i1)) // <- 基底型が同じであれば、型変換が可能
}

/**
 * 名前付きの型と名前のない型
 */
func NamedType() {
    type myInt1 int             // myInt1型の基底型はint型
    type myInt2 myInt1          // myInt2型の基底型はint型 (myInt1型の基底型のint型を引き継ぐ)
    type myArray1 [10]int       // myArray1型の基底型は[10]int型
    type myArray2 [10]myInt1    // myArray2型の基底型は[10]myInt1型 ([10]int型や[10]myInt1型は名前のない型のためそれ自身が基底型となる)

    type namedType1 [10]int     // 名前付きの型で、基底型は[10]int型
    type namedType2 [10]int     // 名前付きの型で、基底型は[10]int型

    var unnamed [10]int         // 名前のない型の変数
    var named1 namedType1       // 名前付きの型の変数
    var named2 namedType2       // 名前付きの型の変数

    named1  = unnamed           // 名前のない型から名前付きの型には代入可能
    unnamed = named1            // 名前付きの型から名前のない型には代入可能
    unnamed = named2            // 名前付きの型から名前のない型には代入可能

    // named1  = named2         // エラーとなる（基底型が同じでも、異なる名前付きの型には代入できない）
    named1 = namedType1(named2) // 基底型が同じであれば、型変換が可能

    var p1 *namedType1          // ポインタ型変数
    var p2 *namedType2          // ポインタ型変数
    p1 = (*namedType1)(p2)      // 型変換が可能
    fmt.Println(p1, p2)
}
