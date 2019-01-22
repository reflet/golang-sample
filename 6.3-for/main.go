package main

import "fmt"

/**
 * 6.3 繰り返し (for文)
 */
func main() {
    // 6.3.1 for文の書き方
    For()

    // 6.3.2 continue文とbreak文
    ContinueBreak()

    // 6.3.2 ラベル付きfor文
    Label()

    // 6.3.3 条件式のみのfor文
    ForLoop()

    // 6.3.4 無限ループとなるfor文
    ForInfiniteLoop()

    // 6.3.5 rangeを用いたfor文
    Range()
}

/**
 * 6.3.1 for文の書き方
 *
 * [for文の構造]
 *
 *  <通常版>
 *  for 初期化文; 条件式; 更新処理文 {
 *      ...
 *  }
 *
 *  <条件式のみ版(while文のような感じ)>
 *  for 条件式 {
 *      ...
 *  }
 *
 *  <無限ループ版>
 *  for {
 *      ...
 *
 *      if 条件式 {
 *          break
 *      }
 *  }
 */
func For() {
    // 通常版
    n := 0
    for i := 1; i <= 10; i++ {
        n += 1
    }
    fmt.Println("通常のfor文: ", n)
}

/**
 * 6.3.2 continue文とbreak文
 */
func ContinueBreak() {
    // continue文の例
    for i := 1; i <= 10; i++ {
        // 偶数であれば出力せず、次の繰り返し処理を実行する
        if (i % 2) == 0 {
            continue
        }
        fmt.Print(i, ",")
    }
    fmt.Println()

    // break文の例
    for i := 1; i <= 10; i++ {
        // iの値が「5」のとき繰り返し処理を終了する
        if i == 5 {
            break
        }
        fmt.Print(i, ",")
    }
    fmt.Println()
}

/**
 * 6.3.2 ラベル付きfor文
 * for文のネストをしている場合、continue文やbreak文の対象をラベルを使うことで柔軟に指定できる
 *
 * [構造]
 *   // ラベル付きfor文
 *   ラベル名: for文
 *
 *   // ラベル指定のcontinue文
 *   continue ラベル名
 *
 *   // ラベル指定のbreak文
 *   break ラベル名
 */
func Label() {
    // ラベル指定
OuterLoop:
    for i := 0; i < 10; i++ {
    InnerLoop:
        for j := 0; j < 10; j++ {
            fmt.Println(i, j)

            if (i == 0) && (j == 1) {
                fmt.Println("continue")
                continue OuterLoop
            }

            if (i == 1) && (j == 1) {
                fmt.Println("break")
                break OuterLoop
            }

            continue InnerLoop
        }
    }
}

/**
 * 6.3.3 条件式のみのfor文
 */
func ForLoop() {
    // 条件式のみ版
    m := 1
    for m <= 1000 {
        m *= 2
    }
    fmt.Println("条件式のみfor文: ", m)
}

/**
 * 6.3.4 無限ループとなるfor文
 */
func ForInfiniteLoop() {
    // 無限ループ版
    r := 1
    for {
        r *= 2
        if r >= 1000 {
            break
        }
    }
    fmt.Println("無限ループfor文: ", m)
}

/**
 * 6.3.5 rangeを用いたfor文
 *
 * [構造]
 *
 *   <変数の宣言をともなう書き方>
 *   ※ 配列などで使用する場合は、変数2つ
 *   for 変数名1, 変数名2 := range 配列など {
 *       ...
 *   }
 *
 *   ※ チャンネルの場合のみ、変数1つ
 *   for 変数名 := range チャンネル {
 *       ...
 *   }
 *
 *   <宣言済みの変数を使用する書き方>
 *   ※ 配列などで使用する場合は、変数2つ
 *   for 変数名1, 変数名2 = range 配列など {
 *       ...
 *   }
 *
 *   ※ チャンネルの場合のみ、変数1つ
 *   for 変数名 = range チャンネル {
 *       ...
 *   }
 *
 *   ※ 各要素を使用しない場合 (変数の記述を省略)
 *   for range 配列など {
 *       ...
 *   }
 *
 */
func Range() {
    // a: 1バイト(UTF-8)
    // あ,い,う: 3バイト(UTF-8)
    s := "aあいう"

    // マルチバイト文字単位で繰り返し処理を実行する
    for i, r := range s {
        // バイト単位の位置, Unicodeコードポイント, 文字列を表示
        fmt.Println(i, r, string(r))
    }
}
