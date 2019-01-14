package workhello

import "fmt"

/**
 * init関数
 * ※ プログラム実行時に最初に実行される
 * ※ 複数設置可能（ただし、実行順番は決まっていないので注意）
 */
func init() {
	fmt.Println("workhello - init1")
}
func init() {
	fmt.Println("workhello - init2")
}

func Hello() {
	fmt.Println("Hello, workspace.")
}
