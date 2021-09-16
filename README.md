# Clean architecture with Go language templates

[![Build](https://github.com/spilebull/bull/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/spilebull/bull/actions/workflows/reviewdog.yml)
[![Test](https://github.com/spilebull/bull/actions/workflows/unit-test.yml/badge.svg)](https://github.com/spilebull/bull/actions/workflows/unit-test.yml)

## 目的

自分の参照用、備忘録、実験用に作成
テンプレートとしても活用

- 日本語：[テンプレートからリポジトリを作成する - GitHub ヘルプ](https://help.github.com/ja/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)

## ディレクトリ構成

```bash
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
│  └── schema           # DBスキーマはここに格納
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

## 認証

- <https://jwt.io/> 参照
- [GateWay 経由 Users 取得](https://lastonemile-gateway-8f6sxn9s.de.gateway.dev/sample/api/v1/users)

## 利用方法

testdata/http 内の File を Rest Client で試す

``` bash
# 実行
make run

# Automatic テスト
make test

# Database 接続
make db-in

# PostgreSQL 使用
psql template -U test
```

## Request Data Validation

- [go-playground/validator/10](https://github.com/go-playground/validator) を使用
- [validator package · pkg.go.dev](https://pkg.go.dev/github.com/go-playground/validator/v10?tab=doc) カスタムする場合は参照

## Database

- `docker-compose.yml` の `POSTGRES_DB`で指定したデータベースを作成

## Database Migration

### 以下方法

[golang-migrate/migrate](https://github.com/golang-migrate/migrate) をインストール

```bash
# Use on Mac
brew install golang-migrate

# Use on Windows
scoop install migrate
```

## Database Migration File

```bash
make mig-cr name={分かりやすい名前}
```

以下のような File が `db/migrations` 以下に作成

- `000001_create_users_table.down.sql`
- `000001_create_users_table.up.sql`

DDL を up.sql に書き、それを打ち消す SQL を down.sql に書く

注：IF (NOT) EXISTS をつけること。

## Database Migration Execution

まずデータベースの接続情報を環境変数に登録しておく

### 接続情報は以下のような形式

```go
dbdriver://username:password@host:port/dbname?option1=true&option2=false
```

### 環境変数設定

```bash
export POSTGRESQL_URL='postgres://test:test@localhost:5432/test?sslmode=disable'
```

```bash
# migration
make mig-up

# rollback
make mig-down

# migration version
make mig-v
```

### CloudSQL のタイムゾーンについてに注意点

CloudSQL はデフォルトのタイムゾーンがロンドンのため、

`ALTER DATABASE {DB名} SET timezone TO 'Asia/Tokyo';`

を最初に実行する必要がある。

## 使用ライブラリ

> 詳細は `go.mod` を確認。

| 目的 | ライブラリ | ドキュメントリンク|
| ---------- | ------------ |----------------------------------------------------------------------------------------------------------------------- |
| WAF | Chi | [go-chi/chi: lightweight, idiomatic and composable router for building Go HTTP services](https://github.com/go-chi/chi) |
| ORM | facebook/ent | [ent](https://entgo.io/docs/getting-started) |
| Logging   | ZAP | [go-chi/httplog](https://github.com/go-chi/httplog) |
| Errors | errors | [errors - Go 言語](https://xn--go-hh0g6u.com/pkg/errors/) |

## 詳細なドキュメントへのリンク

- [ロギング](doc/logging.md)

## 参考 URL

- [Go はクリーンアーキテクチャの思想を活かせるか？ DMM のゲームプラットフォームに Go 言語を選んだ理由 - ログミー Tech](https://logmi.jp/tech/articles/323451)
- [Go 言語で Clean Architecture を実現して、gomock でテストしてみた - Qiita](https://qiita.com/ogady/items/34aae1b2af3080e0fec4)
