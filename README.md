# GO言語

## ローカル環境構築

```bash
$ docker-compose build
$ docker-compose up -d
$ docker-compose exec go bash
```

## hello world

```bash
# 実行
$ go run hello.go
hello world

# ビルド & 実行
$ go build hello.go
$ ./hello
hello world
```

## ホームページ取得

```bash
$ go run wget.go
200 OK 200 HTTP/2.0 2 0 ...
```

## インポート / 相対パス

```bash
$ go run relative/main.go 
Hello, pkg2.
```

## ワークスペース

```bash
$cd src

# 実行 (コードから直接実行)
$ go run main.go 
workhello - init1
workhello - init2

  ...

00111111001100110011001100110011
false
```

ワークスペースでのコンパイル
```bash
$ go install workpkg/workhello
$ ls -la ../pkg/darwin_amd64/workpkg/
-rw-r--r--  1 onoshoji  staff  3940 Jan 14 09:02 workhello.a
```

main.goのコンパイル & 実行
```bash
$ go build main.go
$ ./main
```

リモートパッケージのインストール
```bash
$ go get github.com/golang/example/stringutil
$ ls -la github.com/golang/example/
drwxr-xr-x  12 onoshoji  staff    384 Jan 14 09:08 .git
-rw-r--r--   1 onoshoji  staff  11358 Jan 14 09:08 LICENSE
-rw-r--r--   1 onoshoji  staff   2634 Jan 14 09:08 README.md
drwxr-xr-x   6 onoshoji  staff    192 Jan 14 09:08 appengine-hello
drwxr-xr-x  16 onoshoji  staff    512 Jan 14 09:08 gotypes
drwxr-xr-x   3 onoshoji  staff     96 Jan 14 09:08 hello
drwxr-xr-x   6 onoshoji  staff    192 Jan 14 09:08 outyet
drwxr-xr-x   4 onoshoji  staff    128 Jan 14 09:08 stringutil
drwxr-xr-x   5 onoshoji  staff    160 Jan 14 09:08 template

$ go run github.com/golang/example/hello/hello.go 
Hello, Go examples!
```

## リテラル

```bash
$ go run literal/main.go
```

## 浮動小数点

```bash
$ go run float/main.go
```

## 演算子

```bash
$ go run operator/main.go
```

## キャスト

```bash
$ go run cast/main.go
```

## 定数

```bash
$ go run constant/main.go
```

## 変数

```bash
$ go run variable/main.go
```

## 条件分岐 (if文)

```bash
$ go run if/main.go
```

## 繰り返し処理 (for文)

```bash
$ go run for/main.go
```

## 条件分岐 (switch文)

```bash
$ go run switch/main.go
```

## ラベルへのジャンプ(goto文)

```bash
$ go run goto/main.go
```

## 関数

```bash
$ go run function/main.go
$ open http://localhost:4000/hello    -> 200 ok
$ open http://localhost:4000/goodbye  -> 200 ok
$ open http://localhost:4000/hoge     -> 404 page not found
```

## 型の宣言

```bash
$ go run type/main.go
```

## 構造体

```bash
$ go run struct/main.go
```

## ポインタ

```bash
$ go run pointer/main.go
```

## new組み込み関数

```bash
$ go run new/main.go
```

## nil

```bash
$ go run nil/main.go
```

## Webサーバ (Helloページ / 入力フォーム)

```bash
$ cd server
$ go run main.go
$ open http://localhost:4000/      <- Helloページ
$ open http://localhost:4000/form  <- 入力フォーム
```

## Webサーバ (電卓アプリ)

```bash
$ cd calculator
$ go run main.go
$ open http://localhost:4000/
```

## Webサーバ (jsonデータの利用)

```bash
$ cd json
$ go run main.go
$ open http://localhost:4000/html
$ open http://localhost:4000/json
```
