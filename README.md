# GO言語

## hello world

```
# 実行
$ go run hello.go
hello world

# ビルド & 実行
$ go build hello.go
$ ./hello
hello world
```

## ホームページ取得

```
$ go run wget.go
&{200 OK 200 HTTP/2.0 2 0 map[Server:[gws] X-Xss-Protection:[1; mode=block] Alt-Svc:[quic=":443"; ma=2592000; v="44,43,39,35"] Expires:[-1] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."] Content-Type:[text/html; charset=Shift_JIS] X-Frame-Options:[SAMEORIGIN] Set-Cookie:[1P_JAR=2019-01-13-23; expires=Tue, 12-Feb-2019 23:55:36 GMT; path=/; domain=.google.co.jp NID=154=JPBbg-zE1Z6Pb36H0Pn3c6ERRWBaoXdZq7I5QLIDQccI3SwbK8xtCDteIsS-Y5cV2vW6binvoeOUs4rJVc1NDTD5RdMU_QX69wo7TrTxwyKbB9YbPQmqaiYYQndwMSSTpOGB6UAkcLynrmvlrgmNASS9wh9-uqE6xbYlCikfiH0; expires=Mon, 15-Jul-2019 23:55:36 GMT; path=/; domain=.google.co.jp; HttpOnly] Date:[Sun, 13 Jan 2019 23:55:36 GMT] Cache-Control:[private, max-age=0]] 0xc000097440 -1 [] false true map[] 0xc000112000 0xc0000d4370} <nil>
~/go
```

## 相対パス

```
$ go run relative/main.go 
Hello, pkg2.
```

## ワークスペース

```
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
```
$ go install workpkg/workhello
$ ls -la ../pkg/darwin_amd64/workpkg/
-rw-r--r--  1 onoshoji  staff  3940 Jan 14 09:02 workhello.a
```

main.goのコンパイル & 実行
```
$ go build main.go
$ ./main
```

リモートパッケージのインストール
``` 
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
