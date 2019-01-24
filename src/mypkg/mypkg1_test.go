package mypkg_test

import(
    "mypkg"
    "testing"
)

/**
 * -------------------------------------------------------
 * 15.5.1 単体テスト (go testコマンド)
 * -------------------------------------------------------
 */

// mypkgパッケージの単体テストプログラム
func TestAdd(t *testing.T) {
    if mypkg.Add(1, 2) != 3 {
        t.Fail()
    }
}


/**
 * -------------------------------------------------------
 * 15.5.2 ベンチマーク (go test --bench コマンド)
 * -------------------------------------------------------
 */

 // MapLoop1関数のベンチマーク
func BenchmarkMapLoop1(b *testing.B) {
    mypkg.MapLoop1()
}

// MapLoop2関数のベンチマーク
func BenchmarkMapLoop2(b *testing.B) {
    mypkg.MapLoop2()
}
