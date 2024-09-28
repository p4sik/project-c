package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func CreateDBClient() *bun.DB {
	var (
		dbUser           = os.Getenv("DB_USER")
		dbPass           = os.Getenv("DB_PASS")
		dbHost           = os.Getenv("DB_HOST")
		dbName           = os.Getenv("DB_NAME")
		dbPort           = os.Getenv("DB_PORT")
		sslMode          = os.Getenv("DB_SSLMODE")
		connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslMode)
	)
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db
}
