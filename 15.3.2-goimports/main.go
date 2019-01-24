package main

import "fmt"

/**
 * -------------------------------------------------------
 * 15.3.2 import文の追加 (goimportsコマンド)
 * -------------------------------------------------------
 * ・ goimportsコマンドのインストール
 *    $ go get golang.org/x/tools/cmd/goimports
 *    ※ goimportsコマンドが追加される ( $GOPATH/bin/goimports )
 *
 * ・ import文の追加処理実行例
 *    $ $GOPATH/bin/goimports main.go
 */
func main() {
	fmt.Println("Hello")
}
