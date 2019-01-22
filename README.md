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

# ビルド & 実行
$ go build hello.go
$ ./hello
```

## ホームページ取得
```bash
$ go run wget.go
```

## インポート / 相対パス
```bash
$ go run 2.3-relative/main.go 
```

## ワークスペース
```bash
# 実行 (コードから直接実行)
$ go run src/main.go 
```

ワークスペースでのコンパイル
```bash
$ go install src/workpkg/workhello
$ ls -la pkg/darwin_amd64/workpkg/
```

main.goのコンパイル & 実行
```bash
$ go build src/main.go
$ src/main
```

## リモートパッケージのインストール
```bash
$ go get github.com/golang/example/stringutil
$ go run github.com/golang/example/hello/hello.go 
```

## 4.2 変数
```bash
$ go run 4.2-variable/main.go
```

## 4.4 リテラル
```bash
$ go run 4.4-literal/main.go
```

## 4.5 定数
```bash
$ go run 4.5constant/main.go
```

## 5.1 演算子
```bash
$ go run 5.1-operator/main.go
```

## 5.4 キャスト
```bash
$ go run 5.4-cast/main.go
```

## 6.1 条件分岐 (if文)
```bash
$ go run 6.1-if/main.go
```

## 6.2 条件分岐 (switch文)
```bash
$ go run 6.2-switch/main.go
```

## 6.3 繰り返し処理 (for文)
```bash
$ go run 6.3-for/main.go
```

## 6.4 ラベルへのジャンプ(goto文)
```bash
$ go run 6.4-goto/main.go
```

## 6.4 浮動小数点
```bash
$ go run 6.4-float/main.go
```

## 7.1 Webサーバの起動とHTMLフォーム
```bash
$ go run 7.1-app-server/main.go
$ open http://192.168.99.100:8080/      <- HelloWorldページ
$ open http://192.168.99.100:8080/form  <- 入力フォーム
```

## 7.2 電卓の実装
```bash
$ go run 7.2-app-calculator/main.go
$ open http://192.168.99.100:8080/
```

## 8 関数
```bash
$ go run 8-function/main.go
$ open http://localhost:4000/hello    -> 200 ok
$ open http://localhost:4000/goodbye  -> 200 ok
$ open http://localhost:4000/hoge     -> 404 page not found
```

## 9.1 型の宣言
```bash
$ go run 9.1-type/main.go
```

## 9.2 構造体
```bash
$ go run 9.2-struct/main.go
```

## 9.3 ポインタ
```bash
$ go run 9.3-pointer/main.go
```

## 9.3 new組み込み関数
```bash
$ go run 9.3-new/main.go
```

## 9.3 nil
```bash
$ go run 9.3-nil/main.go
```

## 9.4 JSONデータの使用したWebアプリケーション
```bash
$ go run app-json/main.go
$ open http://192.168.99.100:8080/html
$ open http://192.168.99.100:8080/json
```

## 10.1 メソッド
```bash
$ go run 10.1-method/main.go
```

## 10.2 インターフェース
```bash
$ go run 10.2-interface/main.go
```

## 10.3 Webアプリケーションのルーティング
```bash
$ go run 10.3-app-routing/main.go
$ open http://192.168.99.100:8080/hello
$ open http://192.168.99.100:8080/goodbye
```

## 11.1 配列
```bash
$ go run 11.1-array/main.go
```

## 11.1 スライス
```bash
$ go run 11.1-slice/main.go
```

## 11.2 マップ
```bash
$ go run 11.2-map/main.go
```

## 12.1 関数の戻り値によるエラー
```bash
$ go run 12.1-error/main.go
```

## 12.2 パニックによるエラー
```bash
$ go run 12.2-panic/main.go
```

## 12.3 HTMLテンプレートエンジンとエラー処理
```bash
$ go run 12.3-html-template/main.go
```

## 13 簡単なWebアプリケーション
```bash
$ go run 13-web-app/main.go
```

## 14.1 ゴルーチン
```bash
$ go run 14.1-goroutine/main.go
```

## 14.2 ゴルーチン間の通信（チャンネル）
```bash
$ go run 14.2-channel/main.go
```

## 14.3 共有変数と排他制御
```bash
$ go run 14.3-shared-variable/main.go
```

## 14.4 ゴルーチンとOSスレッド
```bash
$ go run 14.4-thread/main.go
```

## 15.1 Printf関数
```bash
$ go run 15.1-printf/main.go
```

