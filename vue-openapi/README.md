# Vue3 & Golang で実装する openAPI

## Docs

sample-api.yaml に API の仕様を記述する

## Front

グローバルに swaggerCLI をインストールする
https://www.npmjs.com/package/@apidevtools/swagger-cli

下記コマンドで docs に定義していた swagger 形式の yaml で定義した api の仕様書を出力する

```
$ swagger-cli bundle ../docs/sample-api.yaml --outfile sample-api.gen.yaml --type yaml
```

VueCreate したら openapi-generator-cli をインストール
https://github.com/OpenAPITools/openapi-generator-cli

下記コマンドで src 内に定義した API 仕様を TypeScript のコードとして生成する

```
$ openapi-generator-cli generate -g typescript-axios -i ../docs/sample-api.yaml -o ./src/generated/
```

prism でモックを扱う
https://github.com/stoplightio/prism
