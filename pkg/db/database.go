package database

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	DefaultMaxOpenConnectionsDev  = 20
	DefaultMaxOpenConnectionsProd = 50

	DefaultMaxOpenConnections = DefaultMaxOpenConnectionsDev
)

type DB interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type TxFactory interface {
	Transaction(isolation sql.IsolationLevel, readOnly bool) (*Tx, error)
	TransactionContext(context context.Context, isolation sql.IsolationLevel, readOnly bool) (*Tx, error)
}

type Scanner interface {
	Scan(dest ...any) error
}

type Driver interface {
	DB
	TxFactory
	SetIfNil(tx *DB)

	ApplyMigrations(fs fs.FS) error
}

type Option func(sql *sql.DB)

func WithMaxOpenConnections(n int) Option {
	return func(sql *sql.DB) {
		sql.SetMaxOpenConns(n)
	}
}

func WithMaxIdleConnections(n int) Option {
	return func(sql *sql.DB) {
		sql.SetMaxIdleConns(n)
	}
}

func WithConnectionMaxIdleTime(d time.Duration) Option {
	return func(sql *sql.DB) {
		sql.SetConnMaxIdleTime(d)
	}
}

func WithConnectionMaxLifetime(d time.Duration) Option {
	return func(sql *sql.DB) {
		sql.SetConnMaxLifetime(d)
	}
}

type Database struct {
	*sql.DB
	logger *slog.Logger
}

func NewDatabase(logger *slog.Logger, dataSource string, options ...Option) *Database {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(DefaultMaxOpenConnections)

	for _, opt := range options {
		opt(db)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Database{
		DB:     db,
		logger: logger,
	}
}

func (db *Database) Transaction(isolation sql.IsolationLevel, readOnly bool) (*Tx, error) {
	return db.TransactionContext(context.Background(), isolation, readOnly)
}

func (db *Database) TransactionContext(context context.Context, isolation sql.IsolationLevel, readOnly bool) (*Tx, error) {
	tx, err := db.DB.BeginTx(context, &sql.TxOptions{
		Isolation: isolation,
		ReadOnly:  readOnly,
	})
	if err != nil {
		return nil, err
	}
	return &Tx{
		Tx:     tx,
		logger: db.logger,
	}, nil
}

func (db *Database) ApplyMigrations(fs fs.FS) error {
	goose.SetBaseFS(fs)
	goose.SetLogger(&gooseLogger{
		Logger: db.logger,
	})
	err := goose.Up(db.DB, "migrations")
	if err != nil {
		db.logger.Error(err.Error())
	}
	return err
}

func (db *Database) SetIfNil(tx *DB) {
	if *tx == nil {
		*tx = db
	}
}

// Tx transaction
type Tx struct {
	*sql.Tx
	logger *slog.Logger
}

func (tx *Tx) CommitOrRollback(err *error) {
	if err != nil && *err != nil {
		if e := tx.Rollback(); e != nil {
			tx.logger.Error(fmt.Sprintf("rollback transaction error: %s", e))
		}
	} else {
		if e := tx.Commit(); e != nil {
			tx.logger.Error(fmt.Sprintf("commit transaction error: %s", e))
		}
	}
}

func ScanRows[T any](rows *sql.Rows, scanRow func(Scanner) (*T, error)) ([]T, error) {
	var result []T
	for rows.Next() {
		value, err := scanRow(rows)
		if err != nil {
			return result, err
		}
		result = append(result, *value)
	}
	return result, rows.Err()
}
