package main

import (
    "encoding/json"
    "fmt"
    "strconv"
)

func main() {
    // 12.1.1 errorインターフェース
    example11()

    // 12.1.2 型アサーションによるエラー処理
    example12()
}

/**
 * -------------------------------------------------------
 * 12.1.1 errorインターフェース
 * -------------------------------------------------------
 */
func example11() {
    i, err := strconv.Atoi("10")
    if err != nil {
        fmt.Println("[エラー出力]", err)
        return
    }
    fmt.Println(i + 1) // 11を出力

    // iに0、errにエラー値が設定される
    i, err = strconv.Atoi("あ")
    if err != nil {
        // 数値に変換できないのでエラー出力
        fmt.Println("[エラー出力]", err)
        return
    }
    fmt.Println(i + 1) // エラーのため、この行まで到達しない
}

/**
 * -------------------------------------------------------
 * 12.1.2 型アサーションによるエラー処理
 * -------------------------------------------------------
 * ※ 型アサーション: インターフェース型の値.(型)
 *     -> 型変換の一種 (インターフェース型の内部に格納されている元の型に変換する)
 */

// JSONデータに変換する構造体
type JsonData struct {
    A string
    B int
}

func unmarshal(jsonText string) {
    // JSON形式の文字列を構造体に変換する
    var data JsonData
    err := json.Unmarshal(([]byte)(jsonText), &data)       // 戻り値の型: errorインターフェース

    // JSON文法エラー
    if syntaxErr, ok := err.(*json.SyntaxError); ok {      // errorインターフェースの型アサーションで変換チェック
        fmt.Println("[文法エラー]", syntaxErr)
        fmt.Println("  Offset:", syntaxErr.Offset)
        return
    }

    // JSON型エラー
    if typeErr, ok := err.(*json.UnmarshalTypeError); ok { // errorインターフェースの型アサーションで変換チェック
        fmt.Println("[文法エラー]", typeErr)
        fmt.Println("  Offset:", typeErr.Offset)
        fmt.Println("  Value:", typeErr.Value)
        fmt.Println("  Type:", typeErr.Type)
        return
    }

    // その他のエラー
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("[変換OK]")
    fmt.Println(" ", data)
}

func example12() {
    // JSON形式の文字列を構造体に変換する
    unmarshal(`{"A":"x", "B":1}`) // 正常
    unmarshal(`{"A":"X", "B":1]`) // 文法エラー
    unmarshal(`{"A":2, "B":1}`)   // 型エラー
}
