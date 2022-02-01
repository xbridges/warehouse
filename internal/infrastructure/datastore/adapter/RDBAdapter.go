package adapter

import(
	"context"
	"database/sql"
	"time"

	"github.com/xbridges/warehouse/internal/infrastructure/datastore/driver"
)

type RDBAdapter interface {
	Open(connStr string) error
	Begin() (*sql.Tx, error)
	Close() error
	Conn(ctx context.Context) (*sql.Conn, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping() error
	PingContext(ctx context.Context) error
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Stats() sql.DBStats

	NewWhereCaluses() driver.WhereCaluses
	NewOrderCaluses() driver.OrderCaluses

}

func NewRDBAdapter() RDBAdapter {
	return &driver.PostgresqlDriver{}
}