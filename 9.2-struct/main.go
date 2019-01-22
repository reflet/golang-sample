package main

import (
    "fmt"
    "reflect"
)

func main() {
    // 9.2.1 構造体の書き方
    Struct()

    // 9.2.2 匿名フィールド (埋め込みフィールド)
    AnonymousField()

    // 9.2.3 タグとリフレクション
    TagAndReflection()
}

/**
 * 9.2.1 構造体の書き方
 *
 * [構造]
 *    struct {
 *        フィールド名1 型1
 *        フィールド名2, フィールド名3 型2
 *        ...
 *    }
 */
type point struct {
    name string
    x, y int64
}
func Struct() {
    // 各フィールドを指定して代入
    var p1 point
    p1.name = "point"
    p1.x = 1
    p1.y = 2
    fmt.Println("p1: ", p1.name, p1.x, p1.y)

    // 構造体ごとの設定
    p2 := p1
    p1.name = "この変更は、p2に反映されない"
    fmt.Println("p1:", p1.name, p1.x, p1.y)
    fmt.Println("p2:", p2.name, p2.x, p2.y)

    // 構造体リテラル (フィールド名の指定なし)
    p3 := point{"point", 1, 2}
    fmt.Println("p3:", p3.name, p3.x, p3.y)

    // 構造体リテラル（フィールド名の指定あり）
    p4 := point{name: "point", y: 2} // xは指定しないため、ゼロ値
    fmt.Println("p4:", p4.name, p4.x, p4.y)

    // 構造体リテラル（フィールドなし）
    p5 := point{}
    fmt.Println("p5:", p5.name, p5.x, p5.y)
}

/**
 * 9.2.2 匿名フィールド (埋め込みフィールド)
 *
 * [構造]
 *    struct {
 *        型名
 *        *型名
 *        ...
 *    }
 */
type point2 struct {
    point
    y, z int64
}
func AnonymousField() {
    var p1 point2
    p1.name = "p1name"  //「p1.point.name = "p1name"」と同じ意味
    p1.x = 1            //「p1.point.x = 1」と同じ意味
    p1.y = 2            //「p1.point.y = 2」と同じ意味
    p1.point.y = 20
    p1.z = 3
    fmt.Println("name:", p1.name, p1.point.name)
    fmt.Println("x:", p1.x, p1.point.x)
    fmt.Println("y:", p1.y, p1.point.y)
    fmt.Println("z:", p1.z)

}

/**
 * 9.2.3 タグとリフレクション
 *
 * [構造]
 *    struct {
 *        フィールド名 型 タグ
 *        ...
 *    }
 */
type point3 struct {
    name string `color:"red" size:"12pt"`
    x, y int64
}
func TagAndReflection() {
    var p1 point3
    field := reflect.TypeOf(p1).Field(0)
    fmt.Println("color:", field.Tag.Get("color"))
    fmt.Println("size:", field.Tag.Get("size"))
}
