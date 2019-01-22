package main

import "fmt"

func main() {
    // 10.2.1 インターフェースの実装
    example21()

    // 10.2.2 インターフェースの埋め込み
    example22()

    // 10.2.3 インターフェース型と「==」演算子
    example23()

    // 10.2.4 型アサーション
    example24()

    // 10.2.5 型switch
    example25()
}

/**
 * -------------------------------------------------------
 * 10.2.1 インターフェースの実装
 * -------------------------------------------------------
 */

//「type」キーワードによるインターフェース型の宣言
type myIF interface {
    typeName() string
    add(n int) int
}

// myIFインターフェイスを実装する型の宣言（数値型）
type myInt int

// myInt型のtypeNameメソッドの宣言
func (m myInt) typeName() string {
    return "myInt"
}

// myInt型のaddメソッドの宣言
func (m myInt) add(n int) int {
    return int(m) + n
}

// myInt型のsubメソッドの宣言
func (m myInt) sub(n int) int {
    return int(m) - n
}

// myIFインターフェースを実装する型の宣言（文字列型）
type myString string

// myString型のtypeNameメソッドの宣言
func (m myString) typeName() string {
    return "myString"
}

// myString型のaddメソッドの宣言
func (m myString) add(n int) int {
    return len(m) + n
}

func example21() {
    var n myInt = 1
    var i myIF  = n
    fmt.Println("myInt型-typeName :", i.typeName())           // myInt型のtypeNameメソッドを呼び出す
    fmt.Println("myInt型-typeName :", i.add(2))               // myInt型のaddメソッドを呼び出す
    // i.sub(3)                                               // myIF型からは、subメソッドを呼び出せない

    // ポインタも設定可能
    i = &n
    fmt.Println("myInt型-typeName(ポインタ) :", i.typeName()) // myInt型のtypeNameメソッドを呼び出す
    fmt.Println("myInt型-typeName(ポインタ) :", i.add(2))     // myInt型のaddメソッドを呼び出す

    // myString型の値をmyIF型の変数に設定
    var s myString = "abc"
    i = s
    fmt.Println("myString型-typeName :", i.typeName())        // myString型のtypeNameメソッドを呼び出す
    fmt.Println("myString型-add :", i.add(2))                 // myString型のaddメソッドを呼び出す
}

/**
 * -------------------------------------------------------
 * 10.2.2 インターフェースの埋め込み
 * -------------------------------------------------------
 */

// インターフェース型の宣言
type embedIF interface {
    method1()
}

// インターフェース型の宣言
type myIF2 interface {
    embedIF
    method2()
}

// myIF2インターフェースを実装する型の宣言
type myInt2 int
func (i myInt2) method1() {
    fmt.Println("method1:", i)
}
func (i myInt2) method2() {
    fmt.Println("method2:", i)
}

func example22() {
    var i2 myIF2 = myInt2(1)
    i2.method1()
    i2.method2()
}

/**
 * -------------------------------------------------------
 * 10.2.3 インターフェース型と「==」演算子
 * -------------------------------------------------------
 */
func example23() {
    // 数値1を異なる型(int,byte型)に変換した後に設定
    var i3 interface{} = int(1)
    var i4 interface{} = byte(1)

    // 値が一致していても、型が異なる場合は不一致と判定される
    fmt.Println("i3 == i4:", i3 == i4) // false

    // 異なる型のnilを設定
    var i5 interface{} = nil         // 型指定なしのnil
    var i6 interface{} = (*int)(nil) // *int型のnil

    // 両方ともnilでも、型が異なる場合は不一致と判定
    fmt.Println("i5 == i6:", i5 == i6)

    // 型指定なしのnilとの演算
    fmt.Println("i5 == nil:", i5 == nil) // true
    fmt.Println("i6 == nil:", i6 == nil) // false

    // *int型のnilとの演算
    fmt.Println("i5 == (*int)(nil):", i5 == (*int)(nil)) // false
    fmt.Println("i6 == (*int)(nil):", i6 == (*int)(nil)) // true
}

/**
 * -------------------------------------------------------
 * 10.2.4 型アサーション (型変換)
 * -------------------------------------------------------
 */
func typeAssertion(i interface{}) {
    // 型アサーションでint型に値を変換する
    // ※ iがint型以外の場合はエラー
    n := i.(int)
    fmt.Println("型アサーション:", n + 1)
}
func typeAssertionCheck(i interface{}) {
    // int型の場合のみ、変数okにtrueをセット
    n, ok := i.(int)
    if ok {
        fmt.Println("int:", n + 1)
    }

    // string型の場合のみ、変数okにtrueをセット
    s, ok := i.(string)
    if ok {
        fmt.Println("string:" + s)
    }
}
func example24() {
    typeAssertion(1)
    // typeAssertion("a") // int型に変換できず、実行時エラーとなる

    typeAssertionCheck(1)   //「int: 2」を出力
    typeAssertionCheck("a") //「string: a」を出力
}

/**
 * -------------------------------------------------------
 * 10.2.5 型switch
 * -------------------------------------------------------
 */
func typeSwitch(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("int:", v)      // vはint型
    case string:
        fmt.Println("string:", v)   // vはstring型
    case int32, int64:
        fmt.Println("int32/64:", v) // vはiと同じ型 (interface{}型)
    default:
        fmt.Println("default:", v)  // vはiと同じ型 (interface{}型)
    }
}
func typeSwitchNil(i interface{}) {
    switch v := i.(type) {
    case nil:
        fmt.Println("nil:", v)  // 型のないnilのみ該当する
    case *int:
        fmt.Println("*int:", v) // *int型のnilはこちらに該当する
    }
}
func example25() {
    typeSwitch(1)         //「int:1」を出力
    typeSwitch("a")       //「string:a」を出力
    typeSwitch(int32(1))  //「int32/64:1」を出力
    typeSwitch(1.1)       //「default:1.1」を出力

    typeSwitchNil(nil)    //「nil:<nil>」を出力
    var n24 *int = nil
    typeSwitchNil(n24)    //「*int:<nil>」を出力
}
