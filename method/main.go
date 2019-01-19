package main

import "fmt"

type myInt int
func main() {
    /**
     * ---------------------------------------------------
     * メソッドの宣言（値レシーバ）
     * ---------------------------------------------------
     */
    // myInt型のメソッド呼び出し
    var v myInt = 1
    fmt.Println(v.add(2))

    // myStruct型のメソッド呼び出し
    var s = myStruct{4,5}
    fmt.Println(s.add(6))

    // 型の宣言では、メソッドを引き継がない
    type myStruct2 myStruct
    s2 := myStruct2{4,5}
    fmt.Println(s2)
    // fmt.Println(s2.add(6)) <- addメソッドが無くてエラーとなる


    /**
     * ---------------------------------------------------
     * メソッドの宣言（ポインタレシーバ）
     * ---------------------------------------------------
     */
    // addValue()メソッドは、レシーバがコピーされるため、呼び出し元のレシーバ (変数:v) を変更できない
    var v2 myInt = 1
    v2.addValue(2)
    fmt.Println(v2)

    // addPointer()メソッドは、呼び出し元のレシーバをポインタとして受け取るため、レシーバ (変数:v) を変更可能
    var p2 *myInt = &v2
    p2.addPointer(2)
    fmt.Println(*p2)
    fmt.Println(v2)


    /**
     * ---------------------------------------------------
     * メソッドの呼び出し
     * ---------------------------------------------------
     */
    // ポインタ型の変数から、値レシーバのメソッド呼び出し
    var v3 myInt = 1
    var p3 *myInt = &v3
    p3.methodValue()

    // ポインタ型ではない変数から、ポインタレシーバのメソッドの呼び出し
    v3.methodPointer()

    p3 = nil
    p3.methodPointer()   // nilポインタから、ポイントレシーバのメソッドを呼び出す
    //p3.methodValue()   // nilポインタから、値レシーバのメソッドを呼び出す <- エラー


    /**
     * ---------------------------------------------------
     * メソッド値とメソッド式
     * ---------------------------------------------------
     */
    // メソッド値
    var v4 myInt = 1
    methodValue := v4.method
    methodValue("メソッド値:")

    // 値のメソッド式
    methodExpr := myInt.method
    methodExpr(2, "値のメソッド式:")

    // ポインタのメソッド式
    methodExprPtr := (*myInt).method
    v4 = 3
    methodExprPtr(&v4, "ポインタのメソッド式")


    /**
     * ---------------------------------------------------
     * メソッドと匿名フィールド
     * ---------------------------------------------------
     */
     p3d := point3d{point2d{2,3}, 4}
     fmt.Println(p3d.mulXY())       // point2d型のmulXYメソッドを呼び出す
     fmt.Println(p3d.mul())         // point3d型のmulメソッドを呼び出す
     fmt.Println(p3d.point2d.mul()) // point2d型のmulメソッドを呼び出す
}

/**
 * -------------------------------------------------------
 * メソッドの宣言（値レシーバ）
 * -------------------------------------------------------
 */
type myStruct struct {
    x, y int
}
// myInt型（整数値）のメソッド宣言
func (m myInt) add(n int) myInt {
    return m + myInt(n)
}
// myStruct型（構造体型）のメソッド宣言
func (m myStruct) add(n int) myStruct {
    return myStruct{m.x + n, m.y + n}
}

/**
 * -------------------------------------------------------
 * メソッドの宣言（ポインタレシーバ）
 * -------------------------------------------------------
 */

func (m myInt) addValue(n myInt) {
    m += n
}
func (m *myInt) addPointer(n myInt) {
    *m += n
}

/**
 * -------------------------------------------------------
 * メソッドの呼び出し
 * -------------------------------------------------------
 */
func (m myInt) methodValue() {
    fmt.Println("値レシーバ", m)
}
func (m *myInt) methodPointer() {
    if m == nil {
        fmt.Println("ポインタレシーバ： nil")
    } else {
        fmt.Println("ポインタレシーバ：", *m)
    }
}

/**
 * -------------------------------------------------------
 * メソッド値とメソッド式
 * -------------------------------------------------------
 */
func (m myInt) method(s string) {
    fmt.Println(s, m)
}

/**
 * -------------------------------------------------------
 * メソッドと匿名フィールド
 * -------------------------------------------------------
 */
type point2d struct {
    x, y int
}
type point3d struct {
    point2d  // 匿名フィールド(mulXYメソッドとmulメソッドを引き継ぐ)
    z int
}
// xとyの乗算
func (p2d point2d) mulXY() int {
    return p2d.x * p2d.y
}
// xとyの乗算(mulXYメソッドと同じ処理)
func (p2d point2d) mul() int {
    return p2d.x * p2d.y
}
// x,y,zの乗算 (point3d特有のメソッド)
func (p3d point3d) mul() int {
    return p3d.x * p3d.y * p3d.z
}
