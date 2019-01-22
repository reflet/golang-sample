package main

import (
    "fmt"
    "net/http"
)

// serverHTML型の宣言
type serverHTML struct {
    title string
    body string
}

// ServeHTTPメソッドの宣言
func (s *serverHTML) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h := `
<html><head></head><body>
` + s.body + `
</body></html>
`
    fmt.Fprint(w, h)
}

/**
 * 10.3 Webアプリケーションのルーティング
 *     (メソッドとインターフェースによるルーティング)
 */
func main() {
    helloHTML   := &serverHTML{"Hello Title", "Hello World"}
    goodbyeHTML := &serverHTML{"Goodbye Title","Goodbye"}
    http.Handle("/hello", helloHTML)
    http.Handle("/goodbye", goodbyeHTML)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
