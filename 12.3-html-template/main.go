package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type serverHTML struct {
    Title string
    Body  string
}

// テンプレートの設定
// ※ テンプレート読込: func ParseFiles(filename ...string) (*Template, error) {}
// ※ エラーチェック: func Must(t *Template, err error) *Template { }
var tmpl = template.Must(template.ParseFiles("12.3-html-template/html.tmpl")) // <- $GOPATHからのパスを指定

// ServeHTTPメソッドの宣言
func (s *serverHTML) ServeHTTP(w http.ResponseWriter, r * http.Request) {
    // HTTPヘッダの設定
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    // テンプレートとデータを組み合わせたHTMLをクライアントへ送信する
    // ※ 送信処理: func (t *Template) Execute(wr io.Writer, data interface{}) error {}
    if err := tmpl.Execute(w, s); err != nil {
        // エラー時はパニック発生
        panic(err)
    }
}

/**
 * -------------------------------------------------------
 * 12.3 HTMテンプレートエンジンとエラー処理
 * -------------------------------------------------------
 *
 * <実行例>
 *  $ go run 12.3-html-template/main.go
 *  $ open http://192.168.99.100:8080/hello
 *  $ open http://192.168.99.100:8080/goodbye
 */
func main() {
    helloHTML := &serverHTML{"Hello Title", "Hello"}
    goodbyeHTML := &serverHTML{"Goodbye Title", "Goodbye"}
    http.Handle("/hello", helloHTML)
    http.Handle("/goodbye", goodbyeHTML)

    // Webサーバ起動
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
