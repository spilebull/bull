# コンテナでのマイグレーション実行

## 手順

- postgres コンテナを立ち上げる
- コンテナ内に DB を作成する
- DB にスキーマを作成する（スキーママイグレーション）

```shell
postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres12 dropdb simple_bank

migrateup:
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
```

## github action で db 統合テスト自動化

[How to setup Github Actions for Go + Postgres to run automated tests - DEV Community 👩‍💻👨‍💻](https://dev.to/techschoolguru/how-to-setup-github-actions-for-go-postgres-to-run-automated-tests-81o)

## postgre サービスを追加
