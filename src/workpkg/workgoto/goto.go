package workgoto

import "fmt"

/**
 * ラベルへのジャンプ(goto文)
 */
func Goto() {
    // goto文でラベルにジャンプする
    goto label

    // 実行されない文
    fmt.Println("この文は表示されません")

label:
    fmt.Println("goto文による移動")
}
