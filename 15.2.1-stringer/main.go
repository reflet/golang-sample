package main

import "fmt"

/**
 * -------------------------------------------------------
 * 15.2.1 列挙値の文字列化（stringerコマンド）
 * -------------------------------------------------------
 * ・ stringerコマンドのインストール
 *    $ go get golang.org/x/tools/cmd/stringer
 *    ※ stringerコマンドが追加される ( $GOPATH/bin/stringer )
 *
 * ・ コード変換処理実行例
 *    $ $GOPATH/bin/stringer -type DayOfWeek main.go
 *    ※ dayofweek_string.goファイルが生成される
 */

type DayOfWeek int

const (
    Sunday DayOfWeek = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    fmt.Println(Sunday, Saturday)
}
