# go-starter
golang学習用です

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

	`$ ls src/github.com/bitly`

	`go-simplejson`

	`$ tree pkg`

	```
	pkg
	└── darwin_amd64
		└── github.com
			├── bitly
			│   └── go-simplejson.a
			└── example
				└── calc.a

	4 directories, 2 files
```