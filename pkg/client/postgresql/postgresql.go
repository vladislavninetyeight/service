package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config *Config) (conn *pgx.Conn, err error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	conn, err = pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
