package postgres

import (
	"context"
	"fmt"
	"garination.com/db/config"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Queries *Queries
	conn    *pgx.Conn
}

// NewConnection return new PGX connection pool
func NewConnection(cfg *config.Postgres) (*Connection, error) {
	connection := &Connection{}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	queries := New(conn)
	connection.Queries = queries
	return connection, nil
}

func (c *Connection) Close() {
	c.conn.Close(context.Background())
}
