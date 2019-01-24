package main

import "fmt"

/**
 * -------------------------------------------------------
 * 15.2.2 コメント上にあるコマンドの実行
 * -------------------------------------------------------
 * ※ stringerコマンドなでで自動生成したコードは、元のコードが変わる度に再実行が必要となる。
 *    しかし、「go generate」コマンドを使用することで、自動生成するコードを何度も実行せずに済む。
 *
 * ※ コード内に「//go:generate」コメントを埋め込むことでコマンドの自動実行化ができます。
 *
 * ※ コード内に埋め込んだ自動生成コマンドを実行する場合は、下記コマンドを実行する
 *    $ go generate main.go
 *
 * [go generateコマンドで使える変数]
 *  |-------------|--------------------------------------|---------|
 *  |     変数    |           説明                       | 値の例  |
 *  |-------------|--------------------------------------|---------|
 *  | $GOFILE     | 変数を記述したファイル名             | main.go |
 *  | $GOLINE     | 変数を記述したファイル内の行番号     | 1       |
 *  | $GOPACKAGE  | 変数を記述したファイルのパッケージ名 | main    |
 *  | $GOARCH     | CPUアーキテクチャ名                  | amd64   |
 *  | $GOOS       | OS名                                 | windows |
 *  | $DOLLAR     | ドル文字                             | $       |
 *  |-------------|--------------------------------------|---------|
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

// ↓自動実行のためgenerateコメントを記載
//go:generate $GOPATH/bin/stringer -type DayOfWeek $GOFILE
func main() {
    fmt.Println(Sunday, Saturday)
}
