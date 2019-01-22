package main

import (
	"workpkg/workhello"
	"pkg1"
)

/**
 * 2.3.2 ワークスペースでのコンパイル
 *       (go install コマンド)
 */
func main() {
	/**
	 * 複数パッケージの構成 ＋ vendoringの挙動確認
	 *「workhello」パッケージの「Hello」関数を使用する
	 */
	workhello.Hello()

	/**
	 * 相対パスでimportする例
	 * ※「pkg1」パッケージの「F1」関数を使用する
	 */
	pkg1.F1()
}
