package pkg2

import "fmt"

/**
 * pkg1/src.goでインポートされる記述があるが、
 * vendorが優先されて実行さないパッケージの例
 */
func F2() {
	fmt.Println("Hello, src-pkg2.")
}
