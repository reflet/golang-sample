package main

import (
    "fmt"
    "sync"
    "time"
)

/**
 * -------------------------------------------------------
 * 14.3 共有変数と排他制御
 * -------------------------------------------------------
 * ※ 複数のゴルーチン間で値を共有する変数と排他制御の仕組みがある
 * ※ 基本的にはチャンネルを推奨する
 * ※ 扱う問題次第では、変数の値を共有したほうが簡潔に記述できる場合もある
 */
func main() {
    // 14.3.1 共有変数
    example31()

    // 14.3.2 排他制御
    example32()      // 排他制御なしで不正な値になる例 (CPUが2つ以上ある環境で実行!!)
    example32Mutex() // syncパッケージのmutex型で排他制御
}

/**
 * -------------------------------------------------------
 * 14.3.1 共有変数
 * -------------------------------------------------------
 * ※ 同一プログラム(プロセス)内のすべてのゴルーチンは、同じメモリ空間を使用
 * ※ 共有変数 -> グローバル変数のように、どのゴルーチンでも登録/更新できる変数
 * ※ グローバル変数以外にも、関数リテラルとgo文を組み合わせることで、ローカル変数を共有変数にできる
 *    下の例では、グローバル変数(n1) / ローカル変数(n2) とも共有変数となる
 *
 * ※ Go言語の変数は、スレッドセーフではない
 *    スレッドセーフ: 複数のゴルーチンなどの並行処理で扱っても問題が発生しないものを意味する
 *
 *    -> 複数のゴルーチン間で変数を安全に読み書きするには、Mutex型による排他制御が必要！！
 */
// 共有変数（グローバル変数）
var n1 = 0

func example31() {
    // 共有変数（ローカル変数）
    var n2 = 0

    // ゴルーチン待ち合わせ用のチャンネル
    ch := make(chan struct{})

    // どちらのゴルーチンでも同じメモリアドレスを持つ
    go func() {
        n1, n2 = 1, 2
        fmt.Println("go:  ", n1, &n1, n2, &n2)
        close(ch)
    }()

    <-ch // go文のゴルーチンが終了するまで待つ
    fmt.Println("main:", n1, &n1, n2, &n2)
}


/**
 * -------------------------------------------------------
 * 14.3.2 排他制御 (未実装例)
 * -------------------------------------------------------
 * ※ 複数のゴルーチン間で共有変数を読み書きする場合、排他制御が必要！！
 * ※ 排他制御をしていないと、変数の書き込み中に別のゴルーチンが変数を参照すると不正な値を取得してしまう
 */

// 共有変数の宣言
var sharedVar string

// 共有変数に値を書き込む関数
func writeSharedVar(s string) {
    time.Sleep(time.Millisecond) // 1ms待つ (書き込み処理と読み込み処理で実行時間に差を出す)
    sharedVar = s
}

// 共有変数を読み取る関数
func readSharedVar() string {
    return sharedVar
}

func example32() {
    fmt.Println("<排他制御の未実装例>")

    s1 := "AB"
    s2 := "C"
    fmt.Println("s1:", s1, len(s1))
    fmt.Println("s2:", s2, len(s2))

    // 書き込み用ゴルーチンを起動
    go func() {
        // このゴルーチンでは、s1, s2を交互に変数に書き込む
        for {
            writeSharedVar(s1)
            writeSharedVar(s2)
        }
    }()

    for {
        // 特定のタイミングで共有変数の値を取り出す
        s3 := readSharedVar()

        // s1と長さが同じ、かつ、先頭の値がs2と同じ
        if len(s1) == len(s3) && s2[0] == s3[0] {
           fmt.Println("※ 不正な値で停止 -> s3:", s3, len(s3)) // 不正な文字列
           break
        }
    }
}


/**
 * -------------------------------------------------------
 * 14.3.2 排他制御 (実装例)
 * -------------------------------------------------------
 * ※ syncパッケージのmutex型を使用して排他制御を行う
 * ※ 排他制御は、Mutex型以外にも、RWMutex型、Once型、WaitGroup型などもある
 *    ( 詳細はsyncパッケージのドキュメントを参照 )
 */

// Mutex変数の宣言
var mu sync.Mutex

// 共有変数の宣言
var sharedVarMutex string

// 共有変数に値を書き込む関数
func writeSharedVarMutex(s string) {
    mu.Lock()
    defer mu.Unlock()
    time.Sleep(time.Millisecond) // 1ms待つ (書き込み処理と読み込み処理で実行時間に差を出す)
    sharedVarMutex = s
}

// 共有変数を読み取る関数
func readSharedVarMutex() string {
    mu.Lock()
    defer mu.Unlock()
    return sharedVarMutex
}

func example32Mutex() {
    fmt.Println("<排他制御の実装例(Mutex)>")

    s1 := "AB"
    s2 := "C"
    fmt.Println("s1:", s1, len(s1))
    fmt.Println("s2:", s2, len(s2))

    // 書き込み用ゴルーチンを起動
    go func() {
        // このゴルーチンでは、s1, s2を交互に変数に書き込む
        for {
            writeSharedVarMutex(s1)
            writeSharedVarMutex(s2)
        }
    }()

    for {
        // 特定のタイミングで共有変数の値を取り出す
        s3 := readSharedVarMutex()

        // s1と長さが同じ、かつ、先頭の値がs2と同じ
        if len(s1) == len(s3) && s2[0] == s3[0] {
            fmt.Println("s3:", s3, len(s3)) // 不正な文字列
            break
        }
    }
}
