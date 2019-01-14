package main

import "fmt"

func main() {
    Switch()
    Fallthrough()
    Break()
}

/**
 * 条件分岐(switch文)
 *
 * [switch文の構造]
 *
 *  <パターン1>
 *  switch 初期化文; 式1 {
 *      case 式2:
 *         ... (式1の結果が、式2の結果と同じ場合に実行される)
 *
 *      case 式3:
 *         ... (式1の結果が、式2の結果と異なり、式3と同じ場合に実行される)
 *
 *      case 式4, 式5:
 *         ... (式1の結果が、式2,式3の結果と異なり、式4, 式5の結果のどちらかと同じ場合に実行される)
 *
 *      default:
 *         ...
 *  }
 *
 *  <パターン2>
 *  switch {
 *      case 式1:
 *         ... (式1の結果がtrueの場合に実行される)
 *
 *      case 式2:
 *         ... (式1の結果がfalseで、式2の結果がtrueの場合に実行される)
 *
 *      case 式3, 式4:
 *         ... (式1, 式2の結果がfalseで、式3, 式4の結果のどちらかがtrueの場合に実行される)
 *
 *      default:
 *         ...
 *  }
 */
func Switch() {
    x := 1

    // 初期化文で宣言した変数yは、switch文の波カッコ({ })の中だけ有効
    switch y := x + 1; y {
    case 1:
        // 変数zは、このcase文の中だけで有効
        z := x + y
        fmt.Println("case 1:", x, y, z)

    case 2:
        // 変数zは、ここでは利用できない
        fmt.Println("case 2:", x, y)

    } // 変数yはここまで有効
}

/**
 * 条件分岐(switch文 - fallthrough文)
 * ※ そのまま次のcase文も実行したいときに利用する
 */
func Fallthrough() {
    x := 1

    switch x {
    case 1:
        fmt.Println("A")

        // fallthrough文により、次のcase文実行される
        fallthrough

    case 2:
        fmt.Println("B")
        // fallthrough文がないので、switch文を抜ける

    case 3:
        fmt.Println("C")
    }
}

/**
 * 条件分岐(switch文 - break文)
 * ※ case文の途中で抜ける場合に利用する
 */
func Break() {
    x := 1
    y := 2

    switch x {
    case 1:
        fmt.Println("A")

        if y == 2 {
            // switch文から抜ける
            fmt.Println("Break")
            break
        }

        fmt.Println("B")
    }
}
