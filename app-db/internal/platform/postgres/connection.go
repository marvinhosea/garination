package postgres

import (
	"car-marketplace-backend/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Queries *Queries
	conn    *pgx.Conn
}

// NewConnection return new PGX connection pool
func NewConnection(cfg *config.Postgres) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	connection := Connection{}
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	connection.conn = conn

	tx, err := conn.Begin(context.Background())

	if err != nil {
		return nil, err
	}

	connection.Queries = New(tx)

	return conn, nil
}

func (c *Connection) Close() {
	c.conn.Close(context.Background())
}
