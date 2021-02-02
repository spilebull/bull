# sqlc 実行手順

## 手順

- インストール

```
brew install sqlc
```

- `db/migrations/`の下に、migration sql 文を用意する

- `db/query/`の下に、.sql ファイルを用意する。中に CRUD 操作 sql 文を書く

  - `-- name: GetUsers :many`の様なコメントが必要なので、例を見て参考してください。
  - 参考 URL: https://github.com/kyleconroy/sqlc
  - sql 整形ツール　　
    https://sqlfum.pt/

- 実行コマンド

```shell
make sqlcgen
```

- adapter/database/の下にファイルが自動生成された

```
db.go ----　database/sqlのdb interface
models.go ---- migrationから生成されたのentity
querier.go ---- CRUD操作interface
***.sql.go ---- mockファイル
```

- トランザクションを実装する

`store.go`を追加。

- adapter/repository/の用例

```go
type UserRepository struct {
	Store database.Store
}
func (repo *UserRepository) GetUsers(ctx context.Context) (users []domain.User, err error) {
	getUsersRow, err := repo.Store.GetUsers(ctx)
    ...
}
```
