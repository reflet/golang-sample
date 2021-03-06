package main

import "fmt"

/**
 * 6.1 条件分岐(if文)
 * ※ 初期化文が使える
 *
 * [if文の構造]
 *  if 初期化文; 条件式 {
 *     ...
 *  } else if 初期化文; 条件式 {
 *     ...
 *  } else {
 *     ...
 *  }
 */
func main() {
	// 初期化文で利用できるもの
	// 「:=」による変数宣言        if v := 1; 条件式 { ... }
    // 代入文                    if v = 1; 条件式 { ... }
    // インクリメント/デクリメント  if v++; 条件式 { ... }
    // 関数、メソッドの呼び出し     if f(); 条件式 { ... }
    // チャンネル送信             if ch <- 1; 条件式 { ... }
    // チャンネル受信             if <-ch; 条件式 { ... }

    // 初期化文で定義したx変数は、if文の波カッコ({ })内で有効
    if x := 3; x == 1 {
        // xが「1」の場合に出力
        fmt.Println( "x == 1")
    } else if x == 2 {
        // xが「2」の場合に出力
        fmt.Println( "x == 2")
    } else {
        // すべてが異なる場合に出力
        fmt.Println(x)
    }
}
