# go-starter
golang学習用です

## 環境変数GOPATH設定

- `~/.zshrc`に追記

	```
	export GOPATH=$HOME/go-starter #goプログラムソースの置き場
	export PATH=$PATH:$GOPATH/bin #go installしたコマンドに常にパスを通す
	```

## Goのプロジェクト構成とパッケージ

	```
	$ tree go-starter
	go-starter
	├── bin # go install時の格納先
	├── pkg # 依存パッケージオブジェクトファイル
	└── src # プログラムソースコード
	```

## ビルドと実行
- 実行

	```
	$ cd $GOPATH/src/exam
	$ go run main.go
	Hello, World!
	```

- ビルド

	```
	$ cd $GOPATH/src/exam
	$ go install
	```

	ビルドした後のプロジェクト
	```
	$ tree go-starter
	go-starter
	├── bin
	│    └── main
	├── pkg
	│    └── darwin_amd64
	│    └── gosample.a
	└── src
	     └── gosample
	          └── gosample.go
	     └── main
	          └── main.go
	```

## ユニットテストを書いてみる

- calc/add.goをテストするために `add_test.go`を作る
- 実行する

	`$ go test calc`

	`ok  	calc	0.006s`

- エラーを起こさせてみる

	`$ go test calc`

	```
	--- FAIL: TestAdd (0.00s)
	add_test.go:10: Add(1, 2) = 3, want: 4
	FAIL
	FAIL	calc	0.007s
	```

## サード・パーティ製のライブラリーを利用する

- go getコマンドでインストール

	`$ go get github.com/bitly/go-simplejson`

- インストールされていることを確認

	```
	$ ls src/github.com/bitly
	go-simplejson
	```

	```
	$ tree pkg
	pkg
	└── darwin_amd64
	    └── github.com
	    │	     └── bitly
	    │              └── go-simplejson.a
	    └── calc.a

	```