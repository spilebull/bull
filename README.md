# Clean architecture with Go language templates

[![Test](https://github.com/spilebull/bull/actions/workflows/unit-test.yml/badge.svg)](https://github.com/spilebull/bull/actions/workflows/unit-test.yml)
# 実行

`make run`

`testdata/http`内のファイルを`Rest Client`で実行して試してください

> webapi 開発のための Go テンプレート

## ゲートウェイ経由 ユーザーリスト取得パス (認証必須)

> GET https://lastonemile-gateway-8f6sxn9s.de.gateway.dev/sample/api/v1/users

認証方法

- https://github.com/xxx/jwtgen を利用してください

## 目的

GitHub の新しいプロジェクトを作るためのテンプレート

- 日本語：[テンプレートからリポジトリを作成する - GitHub ヘルプ](https://help.github.com/ja/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)

## ディレクトリ構成

```shell
.
├── adapter
│  ├── controller
│  │  └── user.go
│  └── repository
│     └── user.go
├── db
│  └── migrations       # dbスキーママイグレーション
├── doc                 # ドキュメントなど
├── docker-compose.yml
├── Dockerfile
├── domain
│  └── user.go
├── ent
│  └── schema           # ここにDBスキーマを書いてください
├── go.mod
├── go.sum
├── go.test.sh
├── infrastructure
│  ├── httphandler.go
│  ├── postgres.go
│  ├── router.go
│  └── server.go
├── main.go
├── Makefile
├── middleware
├── README.md
├── reference
│  └── user
│     └── schema.yml
├── testdata
│  └── http             # 手動テスト
├── testutil
│  └── testutil.go
├── tmp
│  ├── build-errors.log
│  └── main
└── usecase
   ├── request.go
   └── user.go
```

## 利用方法

### 自動テスト

`make test`

## db 接続

`make db-in`
`psql template -U test`

## リクエストデータのバリデーション（検証）

[go-playground/validator/10](https://github.com/go-playground/validator)を利用している。

まずは[validator package · pkg.go.dev](https://pkg.go.dev/github.com/go-playground/validator/v10?tab=doc)で目的にあったルールを探す

## データベース

`docker-compose.yml`で`POSTGRES_DB`に指定したデータベースが作成されます。

## データベースマイグレーション

以下の方法で[golang-migrate/migrate](https://github.com/golang-migrate/migrate)をインストールしてください。

Mac の方は`Homebrew`でインストールしてください

`$ brew install golang-migrate`

Windows の方は`scoop`などでインストールしてください

`$ scoop install migrate`

### マイグレーションファイルの作成

`make mig-cr name=わかりやすい名前`

以下のようなファイルが`db/migrations`以下に作成される

- `000001_create_users_table.down.sql`
- `000001_create_users_table.up.sql`

DDL を up.sql に書き、それを打ち消す SQL を down.sql に書く

注：IF (NOT) EXISTS をつけること。

### マイグレーションの実行

まずデータベースの接続情報を環境変数に登録しておく
接続情報は以下のような形式
`dbdriver://username:password@host:port/dbname?option1=true&option2=false`

環境変数設定

`$ export POSTGRESQL_URL='postgres://test:test@localhost:5432/test?sslmode=disable'`

マイグレーション

`make mig-up`

逆マイグレーション

`make mig-down`

マイグレーションのバージョン確認

`make mig-v`

### CloudSQL のタイムゾーンについてに注意点

CloudSQL はデフォルトのタイムゾーンがロンドンのため、

`ALTER DATABASE {DB名} SET timezone TO 'Asia/Tokyo';`

を最初に実行する必要がある。

## 使用ライブラリ

> 詳細は`go.mod`を確認。

| 目的       | ライブラリ   | ドキュメントリンク                                                                                                      |
| ---------- | ------------ | ----------------------------------------------------------------------------------------------------------------------- |
| WAF        | Chi          | [go-chi/chi: lightweight, idiomatic and composable router for building Go HTTP services](https://github.com/go-chi/chi) |
| ORM        | facebook/ent | [ent](https://entgo.io/docs/getting-started/)                                                                           |
| ロギング   | ZAP          | [go-chi/httplog](https://github.com/go-chi/httplog)                                                                     |
| エラー処理 | errors       | [errors - Go 言語](https://xn--go-hh0g6u.com/pkg/errors/)                                                               |

## 詳細なドキュメントへのリンク

- [ロギング](doc/logging.md)

## 参考 URL

- [Go はクリーンアーキテクチャの思想を活かせるか？　 DMM のゲームプラットフォームに Go 言語を選んだ理由 - ログミー Tech](https://logmi.jp/tech/articles/323451)
- [Go 言語で Clean Architecture を実現して、gomock でテストしてみた - Qiita](https://qiita.com/ogady/items/34aae1b2af3080e0fec4)
