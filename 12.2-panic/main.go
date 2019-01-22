package main

import (
    "fmt"
)

func main() {
    // 12.2.1 パニックの発生
    // example21()    // <- 同時に実行できないのでコメントアウト

    // 12.2.2 パニックとdefer文
    // example22()    // <- 同時に実行できないのでコメントアウト

    // 12.2.3 パニックからの回復
    example23()

    // 12.2.4 ランタイムパニック
    example24()
}

/**
 * -------------------------------------------------------
 * 12.2.1 パニックの発生
 * -------------------------------------------------------
 * ※ パニックが発生すると残りの処理はスキップされ、実行されません。
 *
 * <実行結果の例>
 *  [main start]
 *  [funcPanic2 start]
 *  [funcPanic1 start]
 *  panic: [funcPanic1 panic]
 *
 *  goroutine 1 [running]:
 *  main.funcPanic1()
 *          /go/error/main.go:98 +0x79
 *  main.funcPanic2()
 *          /go/error/main.go:103 +0x66
 *  main.example21()
 *          /go/error/main.go:108 +0x66
 *  main.main()
 *          /go/error/main.go:12 +0x2a
 *  exit status 2
 *
 */
func funcPanic1() {
    fmt.Println("[funcPanic1 start]")
    panic("[funcPanic1 panic]")       // パニック発生
    fmt.Println("[funcPanic1 end]")   // パニック発生のため、実行されない
}
func funcPanic2() {
    fmt.Println("[funcPanic2 start]")
    funcPanic1()                      // この関数呼び出し内でパニック発生
    fmt.Println("[funcPanic2 end]")   // パニック発生のため、実行されない
}
func example21() {
    fmt.Println("[main start]")
    funcPanic2()                      // この関数呼び出し内でパニック発生
    fmt.Println("[main end]")         // パニック発生のため、実行されない
}

/**
 * -------------------------------------------------------
 * 12.2.2 パニックとdefer文
 * -------------------------------------------------------
 * ※ パニックが発生すると関数の残りの処理は実行されませんが、defer文で登録した処理は例外となり、
 *    通常の関数終了時と同じように、defer文で登録した処理が呼び出されます。
 *
 * <実行結果の例>
 *  [funcPanicDefer1 panic start]
 *  [defer funcPanicDefer1]
 *  [defer funcPanicDefer2]
 *  [defer main]
 *  panic: [funcPanicDefer1 panic]
 *
 *  goroutine 1 [running]:
 *  main.funcPanicDefer1()
 *          /go/error/main.go:167 +0xd5
 *  main.funcPanicDefer2()
 *          /go/error/main.go:171 +0x7e
 *  main.example22()
 *          /go/error/main.go:175 +0x7e
 *  main.main()
 *          /go/error/main.go:20 +0x2a
 *  exit status 2
 *
 */
func funcPanicDefer1() {
    defer fmt.Println("[defer funcPanicDefer1]") // パニック発生中でも実行
    fmt.Println("[funcPanicDefer1 panic start]")
    panic("[funcPanicDefer1 panic]")             // パニック発生
}
func funcPanicDefer2() {
    defer fmt.Println("[defer funcPanicDefer2]") // パニック発生中でも実行
    funcPanicDefer1()
}
func example22() {
    defer fmt.Println("[defer main]")            // パニック発生中でも実行
    funcPanicDefer2()
}

/**
 * -------------------------------------------------------
 * 12.2.3 パニックからの回復
 * -------------------------------------------------------
 * ※ パニックは基本的には異常時のエラーのみで使用するべきで、プログラムを終了するしかない場合に利用します。
 *    回復できるのは、パニック時の動作がドキュメントなどにきさいされている特殊なケースのみ。
 * ※ 通常は、関数の戻り値でエラー処理を行うことをお勧めします。
 *
 * <実行結果の例>
 *  [main start]
 *  [funcRecover start]
 *  [funcPanic start]
 *  [recover] [funcPanic panic]
 *  [main end]
 *
 */
func funcPanic() {
    fmt.Println("[funcPanic start]")
    panic("[funcPanic panic]")       // パニック発生
    fmt.Println("[funcPanic end]")   //  パニック発生のため、実行されない
}
func funcRecover() {
    fmt.Println("[funcRecover start]")

    defer func() {
        // recover関数の呼び出しで通常の処理の流れに戻る
        r := recover()
        if r != nil {
            fmt.Println("[recover]", r)
        }
    }()

    funcPanic()
    fmt.Println("[funcRecover end]") //  パニック発生のため、実行されない
}
func example23() {
    fmt.Println("[main start]")
    funcRecover()
    fmt.Println("[main end]")        // パニックから回復したため、実行される
}


/**
 * -------------------------------------------------------
 * 12.2.4 ランタイムパニック
 * -------------------------------------------------------
 * ※ 実行時エラーのことをいい、nilへのアクセスやゼロ除算などで発生します。
 * ※ panic関数による通常のパニックと同じようにrecover関数で回復できます。
 */
func funcRuntimePanic() {
    var a []int       // 初期値はnil
    fmt.Println(a[0]) // nilへのアクセス(a[0])でランタイムパニック発生
}
func example24() {
    defer func() {
        // recover関数の呼び出しで通常の処理の流れに戻る
        r := recover()
        if r != nil {
            fmt.Println("[recover]", r)
        }
    }()

    funcRuntimePanic()
}
