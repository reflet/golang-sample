package main

import "fmt"

/**
 * nil
 * ※ 指し示す先がないことを意味する値
 */
func main() {
    // ポインタ型の変数にnilを設定
    var x *int = nil

    // ポインタ自体の参照はエラーにならない
    fmt.Println(x)

    /**
     * ここでエラーが発生する
     *
     * <nil>
     * panic: runtime error: invalid memory address or nil pointer dereference
     * [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1091548]
     *
     * goroutine 1 [running]:
     * main.main()
     *           /Users/xxxxxx/go/nil/main.go:xx +0x68
     * exit status 2
     */
    //fmt.Println(*x)
}
