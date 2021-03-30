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

### Golangのホスティング
`vercel`を利用。vercelは`api/index.go`をエントリーポイントとして展開する。  
そのため、通常の`go run <file name>`ではなく`vercel dev`で動作し、`http://localhost:3000/api`でアクセスできる。  
（ただ、今回はAPIしかないので、`vercel.json`で設定変更をして`http://localhost:3000`でアクセスできるようにしている）

#### `vercel.json`
- rewrites：URLを書き換えたい場合は利用する
  - `cleanUrls: true`と`{ "source": "/:match*", "destination": "/api/:match*" }`を併用すると、`vercel dev`では動作しても本番環境で動作しない
- cleanUrls：URL末尾の拡張子を外したい場合に利用する
- trailingSlash：falseの場合、末尾に`/`を入れると、末尾にそれがないURLにリダイレクトする
- regions：東京リージョンを利用する場合は`hnd1`

## Goal
JSON APIサーバー

## Reference
- [GoでJSON APIを書く](http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/)
- [Official Runtimes - Vercel](https://vercel.com/docs/runtimes#official-runtimes/go)
- [Configuration Reference - Vercel](https://vercel.com/docs/configuration)
