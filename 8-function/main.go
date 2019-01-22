package main

import (
    "fmt"
    "net/http"
)

/**
 * 8 関数
 */
func main() {
    // 関数を呼び出す
    funcName1()

    // 関数を変数に代入して利用する
    funcName2 := funcName1
    funcName2()

    // 引数を持つ関数の呼び出し
    funcParam1(1)
    funcParam2(1, 2)
    funcParam3(1, 2)

    // 可変長引数を持つ関数の呼び出し
    funcVariadic1()
    funcVariadic1("a")
    funcVariadic1("a", "b")

    slice := []string{"a", "b"}
    funcVariadic1(slice ...)

    // 固定長と可変長の引数を持つ関数の呼び出し
    funcVariadic2(10, "a", "b")

    // 戻り値の値を変数にセット
    a := funcResult1()
    fmt.Println(a)

    // 複数の戻り値を受け取る
    a, b := funcResult2()
    fmt.Println(a, b)

    // 関数呼び出しの結果を引数に利用する
    funcResult3(funcResult2())

    // 名前付き戻り値の関数を呼び出す
    fmt.Println("名前付き戻り値: ", funcNamedResult1())
    fmt.Println("名前付き戻り値: ", funcNamedResult2())
    fmt.Println(funcNamedResult3())
    fmt.Println(funcNamedResult4())
    fmt.Println("名前付き戻り値: ", funcNamedResult5())

    // 関数型を利用した関数を呼び出す
    printIf(true)
    printIf(false)

    // 関数リテラル（クロージャ）
    closure()
    closureLocalVariable()
    f := funcResultFunc()  // <- 戻り値として関数を受け取る
    f()                    // <- 受け取った関数を実行する

    // 関数終了時の関数呼び出し(defer文)
    funcDefer()
    funcDefer2()
    funcDefer3()
    funcDefer4()
    fmt.Println("defer文で戻り値を更新: ", funcDeferResult())

    // 関数を用いたルーティング(Webアプリケーションのルーティング)
    // $ open http://localhost:4000/hello
    // $ open http://localhost:4000/goodbye
    routing()
}

/**
 * 関数の宣言
 */
func funcName1() {
    fmt.Println("Hello")
}

/**
 * 引数を持つ関数の宣言
 */
func funcParam1(n int) {
    fmt.Println(n)
}
func funcParam2(n1 int, n2 int) {     // <- 2つの引数の型を個別に宣言
    fmt.Println(n1, n2)
}
func funcParam3(n1, n2 int) {         // <- 2つの引数の型をまとめて宣言
    fmt.Println(n1, n2)
}
func funcVariadic1(slice ...string) { // <- 可変長引数
    for _, s := range slice {
        fmt.Print(s, ", ")
    }
    fmt.Println()
}
func funcVariadic2(n int, slice ...string) { // <- 固定長と可変長の引数を持つ
    fmt.Print(n, ": ")
    funcVariadic1(slice ...)
}

/**
 * 戻り値
 *
 * [構造]
 *   // 戻り値が1つの場合
 *   func 関数名(仮引数) 戻り値の型 {
 *       return 戻り値の値
 *   }
 *
 *   // 戻り値が2つの場合
 *   func 関数名(仮引数) (戻り値の型1, 戻り値の型2) {
 *       return 戻り値の値1, 戻り値の値2
 *   }
 */
func funcResult1() int {
    return 1
}
func funcResult2() (int, int) {
    return 1, 2
}
func funcResult3(x, y int) {
    fmt.Println("戻り値: ", x + y)
}

/**
 * 名前付き戻り値
 *
 * [構造]
 *   // 戻り値が1つの場合
 *   func 関数名(仮引数) (戻り値の名前 型) {
 *       return 戻り値の値
 *   }
 *
 *   // 戻り値が2つの場合
 *   func 関数名(仮引数) (戻り値の名前1 型1, 戻り値の名前2 型2) {
 *       return 戻り値の値1, 戻り値の値2
 *   }
 *
 *   // 戻り値が2つの場合 (型をまとめて宣言)
 *   func 関数名(仮引数) (戻り値の名前1, 戻り値の名前2 型) {
 *       return 戻り値の値1, 戻り値の値2
 *   }
 */
 func funcNamedResult1() (i int) {       // <- 1つの戻り値を持つ
    i = 10
    return // 10が戻り値
 }
func funcNamedResult2() (i int) {        // <- 1つの戻り値を持つ
    i = 10
    return 20 // 20が戻り値(return文の値を優先するため)
}
func funcNamedResult3() (i int, j int) { // <- 複数の戻り値を持つ
    i, j = 10, 20
    return // 10, 20が戻り値
}
func funcNamedResult4() (i, j int) {     // <- 複数の戻り値を持つ (型をまとめて宣言)
    i, j = 10, 20
    return // 10, 20が戻り値
}
func funcNamedResult5() (i int) {
    return // 0が戻り値
}

/**
 * 関数型
 */
func add(x, y int) int {
    return x + y
}
func mul(x, y int) int {
    return x * y
}
func printIf(b bool) {
    // 関数型の変数を宣言
    var funcVar func(int, int) int

    if b {
        funcVar = add
    } else {
        funcVar = mul
    }
    fmt.Println("関数型: ", funcVar(3, 4))
}

/**
 * 関数リテラル (クロージャ)
 * リテラルのため、クロージャ関数の外の変数もスコープが有効
 */
func closure() {
    f := func(i, j int) int {
        return i + j
    }

    n := f(1,2)
    fmt.Println("関数リテラル(クロージャ): ", n)
}
func closureLocalVariable() {
    n := 10

    // 外側の変数nを出力して、その後インクリメントする
    f := func() {
        fmt.Println("関数リテラル(ローカル変数): ", n)
        n++
    }

    f()                                                  // 10
    fmt.Println("関数リテラル(ローカル変数): ", n)  // 11

    n = 20
    f()                                                  // 20
    fmt.Println("関数リテラル(ローカル変数): ", n)  // 21
}
func funcResultFunc() func() {
    n := 10

    // 戻り値となる関数リテラル
    return func() {
        // 外側にあるローカル変数nを出力する
        fmt.Println("関数を戻り値とする関数: ", n)
    }
}

/**
 * ※ 関数終了時の関数呼び出し(defer文)
 */
func funcDefer() {
    fmt.Println("関数実行(funcDefer) - start")

    defer fmt.Println("defer文の関数呼び出し実行")

    fmt.Println("関数実行(funcDefer) - end")
}
func funcDefer2() {
    // defer文は実行した順番とは逆の順番で実行される
    defer fmt.Println("defer文の関数呼び出し順番: ", 1)
    defer fmt.Println("defer文の関数呼び出し順番: ", 2)
    defer fmt.Println("defer文の関数呼び出し順番: ", 3)
    fmt.Println("関数実行(funcDefer2)")
}
func funcDefer3() {
    // メッセージを生成する関数リテラル
    m := "hoge"
    f := func() string {
        fmt.Println("defer文実行時にメッセージが生成される")
        return m
    }

    // deferで呼び出される関数リテラル
    n := 10
    d := func(msg string) {
        fmt.Println("クロージャ実行(defer文で呼び出される)", msg, n)
    }

    // deferで呼び出される関数が登録される (登録されるだけで、実行自体はしない)
    // ※ 実引数(f())は、defer文の実行時(ここ)で実行される（ローカル変数などは、この時点の値となる）
    defer d(f())

    n = 20      // <- deferの実行に反映される
    m = "fuga"  // <- deferの実行には反映されない
    d = func (msg string) { fmt.Println("上書き", "no exec", 0) } // <- deferでは実行されない
    fmt.Println("関数実行(funcDefer3)", n)
}
func funcDefer4() {
    defer func() { fmt.Println("defer文でクロージャを実行する") }()
    fmt.Println("関数実行(funcDefer4)")
}
func funcDeferResult() (result int) {
    defer func() {
        fmt.Println("更新前戻り値", result)

        // 戻り値を上書きする
        result += 2
    }()

    // return文で戻り値を指定 (ただし、defer文で戻り値が上書きされる)
    return 1
}

/**
 * 関数を用いたルーティング(Webアプリケーションのルーティング)
 */
func routing() {
    helloHandle := func(w http.ResponseWriter, r *http.Request) {
        h := `
<html><head><title>Hello</title></head><body>
  Hello
</body></html>
`
        fmt.Fprint(w, h)
    }

    goodbyeHandle := func(w http.ResponseWriter, r *http.Request) {
        h := `
<html><head><title>Goodbye</title></head><body>
  Goodbye
</body></html>
`
        fmt.Fprint(w, h)
    }

    http.HandleFunc("/hello", helloHandle)
    http.HandleFunc("/goodbye", goodbyeHandle)
    if err := http.ListenAndServe(":4000", nil); err != nil {
        fmt.Println(err)
    }
}
