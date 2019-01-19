package main

import (
    "fmt"
    "html"
    "net/http"
)

/**
 * httpリクエストを受け取るメソッド (Helloページ)
 */
func HelloServer(w http.ResponseWriter, r *http.Request) {
    // HTMLの文字列を設定
    h := `
<html><head><title></title></head><body>
  Hello World
</body></html>
`

    // クライアント（ブラウザ）にHTMLを送信
    fmt.Fprint(w, h)
}

/**
 * httpリクエストを受け取るメソッド (入力フォーム)
 */
func FormServer(w http.ResponseWriter, r *http.Request) {
    v := r.FormValue("input_value")
    h := `
<html><head><title></title></head><body>
  <form action="/form" method="post">
    <input type="text" name="input_value">
    <input type="submit" name="送信"><br>
    入力値:` + html.EscapeString(v) + `
  </form>
</body></html>
`

    // クライアント（ブラウザ）にHTMLを送信
    fmt.Fprint(w, h)
}

/**
 * Webサーバ
 * ※ 参考サイト: http://golang.jp/pkg/http
 */
func main() {
    // Webサーバの起動
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/form", FormServer)
    http.ListenAndServe(":8080", nil)
}
