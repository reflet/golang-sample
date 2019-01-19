package main

import (
    "fmt"
    "html"
    "math/big"
    "net/http"
)

/**
 * httpリクエストを受け取るメソッド (入力フォーム)
 */
func FormServer(w http.ResponseWriter, r *http.Request) {
    // フォームの入力値を取得
    left  := r.FormValue("left")   // 左項目（テキストボックスの値）
    right := r.FormValue("right")  // 右項目（テキストボックスの値）
    op    := r.FormValue("op")     // 演算子（ラジオボタンの値）

    // 文字列を整数に変換 (int型版 - 最大19桁まで対応)
    // leftInt , leftErr  := strconv.Atoi(left)
    // rightInt, rightErr := strconv.Atoi(right)
    //
    // 四則演算
    // ※ 変換エラーがなければ、演算子に従って計算する
    // var result string
    // if (leftErr == nil) && (rightErr == nil) {
    //    switch op {
    //    case "add":   // 加算
    //        result = strconv.Itoa(leftInt + rightInt)
    //    case "sub":   // 減算
    //        result = strconv.Itoa(leftInt - rightInt)
    //    case "multi": // 乗算
    //        result = strconv.Itoa(leftInt * rightInt)
    //    case "div":   // 除算
    //        result = strconv.Itoa(leftInt / rightInt)
    //    }
    //}

    // 文字列を整数に変換 (多倍長整数 - 19桁以上の数字にも対応版)
    leftInt  := &big.Int{}
    rightInt := &big.Int{}
    _, leftOK := leftInt.SetString(left, 10)
    _, rightOK := rightInt.SetString(right, 10)

    // 四則演算
    // ※ 変換エラーがなければ、演算子に従って計算する
    var result string
    if leftOK && rightOK {
        resultInt := &big.Int{}
        switch op {
        case "add":   // 加算
            resultInt.Add(leftInt, rightInt)
        case "sub":   // 減算
            resultInt.Sub(leftInt, rightInt)
        case "multi": // 乗算
            resultInt.Mul(leftInt, rightInt)
        case "div":   // 除算
            resultInt.Div(leftInt, rightInt)
        }
        result = resultInt.String()
    }

    h := `
<html><head><title>電卓アプリ</title></head><body>
  <form action="/" method="post">
    左項目：<input type="text" name="left" value="` + html.EscapeString(left) + `"><br>
    右項目：<input type="text" name="right" value="` + html.EscapeString(right) + `"><br>
    演算子：
    <input type="radio" name="op" value="add" checked> ＋
    <input type="radio" name="op" value="sub"> −
    <input type="radio" name="op" value="multi"> ☓
    <input type="radio" name="op" value="div"> ÷<br>
    <input type="submit" name="送信">
    <hr>
    [フォームの入力値]<br>
    左項目：` + html.EscapeString(left) + `<br>
    右項目：` + html.EscapeString(right) + `<br>
    演算子：` + html.EscapeString(op) + `<br>
    演算結果：` + html.EscapeString(result) + `
  </form>
</body></html>
`
    // クライアント（ブラウザ）にHTMLを送信する
    fmt.Fprint(w, h)
}

/**
 * Webサーバ (電卓アプリ)
 */
func main() {
    // Webサーバの起動
    http.HandleFunc("/", FormServer)
    http.ListenAndServe(":8080", nil)
}
