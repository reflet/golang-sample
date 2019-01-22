package main

import (
    "fmt"
    "time"
)

func funcGo() {
    fmt.Println("go 1")

    // 2ミリ秒待つ（mainのゴルーチンが終了するまで待つ）
    time.Sleep(2 * time.Millisecond)

    // mainのゴルーチンが先に終了するため実行されない
    fmt.Println("go 2")
}

/**
 * -------------------------------------------------------
 * 14.1 ゴルーチン
 * -------------------------------------------------------
 */
func main() {
    // 新規にゴルーチンを作成
    go funcGo()

    // 1ミリ秒待つ（"go 1"を出力するまで待つ）
    time.Sleep(time.Millisecond)

    // funcGo関数による"go 1"を出力後に実行する
    fmt.Println("main 1")
}
