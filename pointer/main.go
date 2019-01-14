package main

import (
    "fmt"
    "unsafe"
)

func main() {
    AddressOperator()
    StructLiteral()
    PointerAndStruct()
    PointerAndConvert()
}

/**
 * アドレス演算子と間接演算子
 */
func AddressOperator() {
    // アドレスを取得する変数の宣言
    var v int32 = 100

    // アドレス演算子
    var pointer *int32 = &v     // ポインタ型の変数に設定
    fmt.Println("ポインタの値:", pointer)

    // 間接演算子
    fmt.Println("間接演算子の結果", *pointer)  // 100を出力

    // 変数vの値を更新
    v = 200
    fmt.Println("変数vの更新:", *pointer)      // 200を出力

    // ポインタ側から値を更新
    *pointer = 300
    fmt.Println("ポインタから値を更新:", v)    // 300を出力
}

/**
 * 構造体リテラルの情報
 */
type myStruct struct {
    field int
}
func StructLiteral() {
    // 構造体リテラルにアドレス演算子を使用
    fmt.Println("&myStruct{1}", &myStruct{1})

    // 構造体変数のフィールドにアドレス演算子を使用
    s := myStruct{1}
    fmt.Println("&s.field:", &s.field)

    // 配列アクセスにアドレス演算子を使用
    array := [10]int{}
    fmt.Println("&array[0]:", &array[0])

    // スライスのアクセスにアドレス演算子を使用
    slice := []int{0}
    fmt.Println("&slice[0]:", &slice[0])
}

/**
 * ポインタと構造体
 */
func PointerAndStruct() {
    s1 := myStruct{1}
    s2 := s1            // コピーで代入される
    s2.field = 2
    fmt.Println("別の値:", s1, s2)

    // ポインタを用いると同じメモリ領域が使用される
    s3 := &s1             // アドレス演算子でメモリのアドレスが代入される
    s3.field = 3          //「(*s3).field = 3」と同じ意味
    fmt.Println("同じ値:", s1, *s3)
}

/**
 * ポインタと型変換
 * ※ ポインタ型をuintptr型に変換する場合は、unsafeパッケージのPointer型を間に挟む必要がある
 *
 * < ポインタ型, Pointer型, uintptr型の型変換ルール >
 * ・ unsafe.Pointer型 <-> 任意のポインタ型 (相互に変換可能)
 * ・ unsafe.Pointer型 <-> uintptr型（相互に変換可能）
 * ・ 型変換によって作成されるポインタは安全ではない場合がある
 */
func PointerAndConvert() {
    // ポインタ型からunsafe.Pointer型への変換
    pointerInt1 := new(int)
    unsafePointer1 := unsafe.Pointer(pointerInt1)
    fmt.Println("unsafePointer1:", unsafePointer1)

    // unsafe.Pointer型からuintptr型への変換
    u := uintptr(unsafePointer1)
    u += 2  // uintptr型は演算可能
    fmt.Println("u:", u)

    // uintptr型からunsafe.Pointer型への変換
    unsafePointer1 = unsafe.Pointer(u)
    fmt.Println("unsafePointer1:", unsafePointer1)

    // unsafe.Pointer型から別のポインタ型への変換
    // ※ *int型 -> unsafe.Pointer型 -> *string型 の変換は、メモリ構造が異なるため、不正なアクセスとなることがある
    pointerInt2 := new(int)                        // *int型が生成される
    unsafePointer2 := unsafe.Pointer(pointerInt2)  // *int型 -> unsafe.Pointer型 へ変換
    fmt.Println(pointerInt2, unsafePointer2)
    // pointerString := (*string)(unsafePointer2)  // unsafe.Pointer型 -> *string型 へ変換 (コンパイル時にエラー)
    // fmt.Println(*pointerString)                 // ここでエラーになる

    /**
     * panic: runtime error: invalid memory address or nil pointer dereference
     * [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1052103]
     */
}
