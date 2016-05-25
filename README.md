# go-starter
golang学習用です

## 環境変数GOPATH設定

- `~/.zshrc`に追記

	```
	export GOPATH=$HOME/go-starter #goプログラムソースの置き場
	export PATH=$PATH:$GOPATH/bin #go installしたコマンドに常にパスを通す
	```

## Goのプロジェクト構成とパッケージ

	$ tree go-starter
	go-starter
	├── bin # go install時の格納先
	├── pkg # 依存パッケージオブジェクトファイル
	└── src # プログラムソースコード


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

## 複数の作業スペース($GOPATH)を持ちたい場合

direnvインストール

	$ brew install derenv


シェルの設定ファイルに direnv 用のフックを仕込む

	$ echo 'eval "$(direnv hook zsh)"' >> ~/.zshrc
	$ source ~/.zshrc


作業スペースとなるディレクトリ作成


	$ mkdir -p ~/go-tmp

direnvの管理対象にしたいディレクトリを追加

	$ direnv allow ~/go-tmp


作成した作業スペースに.envrcファイルを追加

	$ cd ~/go-tmp
	$ vi .envrc  # もしくは、direnv edit .
	unset GOPATH  # GOPATHをUNSETする
	layout go     # カレントディレクトリーをGOPATHにする

管理対象にしたディレクトリに移動すると.envrcに設定した環境変数が反映される

	$ cd ~/go-tmp
    direnv: loading .envrc
    direnv: export ~GOPATH ~PATH

- unset GOPATHにより設定済みのgo作業ディレクトリーを解除した場合
	```
	$ echo $PATH
	{$HOME}/go-tmp/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin
	$ echo $GOPATH
	{$HOME}/go-tmp
	```

- すでにGOPATHが設定されており、unset GOPATHしない場合
	```
	$cat .envrc
	layout go
	```
	```
	$ echo $PATH
	{$HOME}/go-tmp/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:{$HOME}/go-starter/bin
	$ echo $GOPATH
	{$HOME}/go-tmp:/{$HOME}/go-starter
	```

go getでパッケージ取得先が{$HOME}/go-starterから{$HOME}/go-tmpへ変わったことが分かる

	$ pwd
	{$HOME}/go-tmp
	$ go get github.com/bitly/go-simplejson
    $ tree src
      src
      └── github.com
          └── bitly
              └── go-simplejson
                  ├── LICENSE
                  ├── README.md
                  ├── simplejson.go
                  ├── simplejson_go10.go
                  ├── simplejson_go10_test.go
                  ├── simplejson_go11.go
                  ├── simplejson_go11_test.go
                  └── simplejson_test.go



