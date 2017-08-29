アラン・ドノバン, ブライアン・カーニハン著『プログラミング言語Go』（丸善）の練習問題を解いていきます。

## Go ファイルの名前と配置場所
`main` 関数を含む Go ファイルは全て `main.go` としています。　各練習問題の `main.go` ファイルの格納場所は、例えば第1章の練習問題1.2については

  `ch1/ex1_2/main.go`
  
となるようにしています。　一部、性能評価（ベンチマーク）の問題では `main_test.go` にテストコードを書いています（実装コードは `main.go` に書いているものもあります）。

## github リポジトリからのソースコードの取得と実行
Linux, Unix 系の OS では説明は不要でしょう。 Mac は使ったことがないのでここでは扱いません。

以下、Windows での設定方法。

#### 環境変数 GOPATH の設定
本 github リポジトリからソースコードを取得したり、ビルドして実行したりするためには、まず環境変数 `GOPATH` を設定します。　ここにはダウンロードしたソースコードやバイナリを格納するディレクトリへのパスを与えます。　基本的にはどこでもいのですが

  `C:\Users\《ユーザー名》\.go`
  
などがいいと思います（環境変数の設定の仕方はネットで検索してください）。　また、環境変数 `PATH` に `;%GOPATH%\bin` を追記しておきましょう（ダウンロードしたソースコードを実行するのに必要。　ただし、複数のパスを `GOPATH` に設定している場合はこれでは上手くいかないので注意）。

#### ソースコードの取得
GOPATH が設定できたら、例えば第1章の練習問題 1.2 のソースコードを取得するには、適当なディレクトリ上で

  `> go get github.com/waman/exercise-go/ch1/ex1_2`
  
とします。　これを実行すると、上記で設定した `GOPATH` 下の `src` ディレクトリから辿っていって、適当な位置に目的のソースコードがあるはずです（ちなみに、本 github リポジトリ内の全てのソースコードがダウンロードされていると思います）。

#### 実行
上記のコマンドを実行した時点で `ex1_2` 下の `main.go` はビルドされているので、このコードを実行したい場合は単に

  `> ex1_2 a bc def`
  
で実行できます（`a bc def` はプログラムに渡される引数）。　リポジトリ内のソースコードは全てダウンロードされてますが、ビルドは指定したものしかされていないので、別の練習問題のコードを実行したい場合は別途 `go get` コマンドを実行する必要があります。