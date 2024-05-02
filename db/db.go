package db

import (
	"os"
	"sync/atomic"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var master struct {
	dbx atomic.Value
}

func Master() *sqlx.DB {
	return master.dbx.Load().(*sqlx.DB)
}

func init() {
	var dbx *sqlx.DB
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	dbx, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	master.dbx.Store(dbx)
}