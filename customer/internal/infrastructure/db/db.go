package db

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Connection struct {
	*pgx.Conn
}

func NewConnection(config Config) Connection {
	conn, err := pgx.Connect(context.Background(), config.GetDSN())
	if err != nil {
		panic(err)
	}
	return Connection{Conn: conn}
}
