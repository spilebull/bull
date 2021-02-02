/*
	Package infrastructure はデータベースやHTTPリクエストとの窓口です。
	WEBフレームワークの設定、データベースとの接続などを行います。
*/
package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"server/ent"

	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*ent.Client, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		err = fmt.Errorf("[infrastructure.NewSQLHandler] failed to open connection to database: %w", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), err
}
