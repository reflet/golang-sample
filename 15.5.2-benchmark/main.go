package main

/**
 * -------------------------------------------------------
 * 15.5.2 ベンチマーク (go test -bench コマンド)
 * -------------------------------------------------------
 * ※ go testコマンドでは、単体テストと同じ仕組みでベンチマーク（性能を表す指標）を取得することができる。
 * ※ 単体テストと異なる点は、関数名が「Benchmark」で始まることと、引数の型が「*testing.B」型になります。
 *
 * ※ 実行結果の例
 *     $ go test -bench . mypkg
 *       goos: linux
 *       goarch: amd64
 *       pkg: mypkg
 *       BenchmarkMapLoop1-2     2000000000               0.12 ns/op
 *       BenchmarkMapLoop2-2     2000000000               0.08 ns/op
 *       PASS
 *       ok      mypkg   4.785s
 */
