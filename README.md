# Learn Golang
昔挫折したので、また試したい

## Review
### Golangとは
2009年にGoogleの3人のエンジニアによって設計されたプログラミング言語。C++のようなパフォーマンスとセキュリティ、Pythonのような動的言語の速度を併せ持つとされている。

### Golangの依存関係管理
Go Modulesという仕組みがGolang v1.11から導入されており、まず以下のコマンドを入力する。

``` bash
go mod init
```

上記のコマンドを入力することで依存関係を管理する`go.mod`ファイルを作成する。  
その後、使用する依存関係をコードに入力し、`go build <Goファイル名>`とすることで、勝手に依存関係をダウンロードし、また`go.mod`と`go.sum`に依存関係を記録してくれる。

ちなみに、何度も依存関係を解決していると不要なパッケージも溜まるため、`go mod tidy`で削除する。

## Goal
JSON APIサーバー

## Reference
- [GoでJSON APIを書く](http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/)